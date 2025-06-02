require 'net/http'
require 'json'

class LocationController < ApplicationController
  def whereami
    user_info = {
      ip: get_client_ip,
      country: nil,
      language: get_preferred_language
    }

    # Get country code from IP address
    if user_info[:ip]
      country_code = get_country_from_ip(user_info[:ip])
      user_info[:country] = country_code
    end

    render json: user_info, status: 200
  end

  private

  def get_client_ip
    # Try to get the real IP address from various headers (order matters)
    ip = request.env['HTTP_X_FORWARDED_FOR'] || 
         request.env['HTTP_X_REAL_IP'] || 
         request.env['HTTP_CF_CONNECTING_IP'] || 
         request.env['HTTP_X_FORWARDED'] ||
         request.env['HTTP_FORWARDED_FOR'] ||
         request.env['HTTP_FORWARDED'] ||
         request.remote_ip

    # Clean up the IP if it contains multiple addresses (take the first one)
    if ip&.include?(',')
      ip = ip.split(',').first.strip
    end
    
    # Remove any whitespace
    ip = ip&.strip
    
    # Only filter out obvious localhost/private IPs, but be less aggressive
    # Allow all valid public IPs through
    if ip == '127.0.0.1' || ip == '::1'
      return nil  # localhost
    end
    
    # Only filter private network ranges
    if ip&.start_with?('192.168.', '10.')
      return nil  # private networks
    end
    
    # Handle 172.16-31.x.x private range more precisely
    if ip&.start_with?('172.')
      parts = ip.split('.')
      if parts.length >= 2
        second_octet = parts[1].to_i
        if second_octet >= 16 && second_octet <= 31
          return nil  # 172.16.0.0 - 172.31.255.255 private range
        end
      end
    end
    
    ip
  end

  def get_country_from_ip(ip_address)
    return nil unless ip_address

    begin
      # Using HackerRank's API endpoint as specified in requirements
      uri = URI("https://jsonmock.hackerrank.com/api/ip/#{ip_address}")
      response = Net::HTTP.get_response(uri)
      
      if response.code == '200'
        data = JSON.parse(response.body)
        # HackerRank API might have different response format, handle both cases
        if data['country']
          return data['country'] # Direct country field
        elsif data['countryCode']
          return data['countryCode'] # Country code field
        elsif data['status'] == 'success' && data['countryCode']
          return data['countryCode'] # ip-api format fallback
        end
      end
    rescue => e
      Rails.logger.error "Error fetching country from IP #{ip_address}: #{e.message}" if defined?(Rails)
      # Log the actual response for debugging
      Rails.logger.error "Response body: #{response&.body}" if defined?(Rails) && response
    end

    nil
  end

  def get_preferred_language
    accept_language = request.env['HTTP_ACCEPT_LANGUAGE']
    return 'en' unless accept_language

    accepted = accept_language.split(',').map { |language| language.strip.split(";") }
    
    # First try to find a language with hyphen (region code)
    with_region = accepted.find { |lang| lang[0].include?('-') }
    return with_region[0] if with_region
    
    # Otherwise return the first language
    accepted.first&.first || 'en'
  end
end 

