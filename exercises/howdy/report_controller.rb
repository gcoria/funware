class ReportController < ApplicationController
  def create
    report = Report.new(report_params)
    
    if report.save
      render json: report, status: 201
    else
      render json: report.errors, status: 422
    end
  end

  def index
    reports = Report.all.order(:id)
    
    # Filter by date range
    if params[:date_from].present?
      reports = reports.where('date >= ?', Date.parse(params[:date_from]))
    end
    
    if params[:date_to].present?
      reports = reports.where('date <= ?', Date.parse(params[:date_to]))
    end
    
    # Filter by age
    if params[:age_from].present?
      reports = reports.where('age >= ?', params[:age_from].to_i)
    end
    
    # Filter by distance if latitude, longitude, and distance are all provided
    if params[:latitude].present? && params[:longitude].present? && params[:distance].present?
      target_lat = params[:latitude].to_f
      target_lng = params[:longitude].to_f
      max_distance = params[:distance].to_f
      
      reports = reports.select do |report|
        next unless report.latitude.present? && report.longitude.present?
        
        distance = coordinates_distance(target_lat, target_lng, report.latitude, report.longitude)
        distance <= max_distance
      end
    end
    
    render json: reports
  end

  def show
    report = Report.find(params[:id])
    render json: report
  rescue ActiveRecord::RecordNotFound
    render json: { error: 'Report not found' }, status: 404
  end

  private

  def report_params
    params.permit(:date, :name, :gender, :age, :city, :state, :county, :latitude, :longitude)
  end

  # Helper method to calculate distance between coordinates
  def coordinates_distance(lat1, lng1, lat2, lng2)
    # This assumes you have this method available
    # If not implemented, here's a basic haversine formula implementation:
    rad_per_deg = Math::PI / 180  # PI / 180
    rkm = 6371                    # Earth radius in kilometers
    rm = rkm * 1000               # Radius in meters

    dlat_rad = (lat2 - lat1) * rad_per_deg  # Delta, converted to rad
    dlng_rad = (lng2 - lng1) * rad_per_deg

    lat1_rad = lat1 * rad_per_deg
    lat2_rad = lat2 * rad_per_deg

    a = Math.sin(dlat_rad / 2)**2 + Math.cos(lat1_rad) * Math.cos(lat2_rad) * Math.sin(dlng_rad / 2)**2
    c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))

    rkm * c # Delta in kilometers
  end
end
