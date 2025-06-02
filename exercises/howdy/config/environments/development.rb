require "active_support/core_ext/integer/time"

Rails.application.configure do
  # Settings specified here will take precedence over those in config/application.rb.

  # In the development environment your application's code is reloaded any time
  # it changes. This slows down response time but is perfect for development
  # since you don't have to restart the web server when you make code changes.
  config.cache_classes = false

  # Do not eager load code on boot.
  config.eager_load = false

  # Show full error reports.
  config.consider_all_requests_local = true

  # Enable/disable caching. By default caching is disabled.
  # Run rails dev:cache to toggle caching.
  if Rails.root.join('tmp', 'caching-dev.txt').exist?
    config.cache_store = :memory_store
    config.public_file_server.headers = {
      'Cache-Control' => "public, max-age=#{2.days.to_i}"
    }
  else
    config.action_controller.perform_caching = false
    config.cache_store = :null_store
  end

  # Print deprecation notices to the Rails logger.
  config.active_support.deprecation = :log

  # Raise exceptions for disallowed deprecations.
  config.active_support.disallowed_deprecation = :raise

  # Tell Active Support which deprecation messages to disallow.
  config.active_support.disallowed_deprecation_warnings = []

  # Raise an error on page load if there are pending migrations.
  config.active_record.migration_error = :page_load if defined?(ActiveRecord)

  # Highlight code that triggered database queries in logs.
  config.active_record.verbose_query_logs = true if defined?(ActiveRecord)

  # Debug mode disables concatenation and preprocessing of assets.
  # This option may cause significant delays in view rendering with a large
  # number of complex assets.
  config.assets.debug = true if defined?(ActionCable)

  # Suppress logger output for asset requests.
  config.assets.quiet = true if defined?(ActionCable)

  # DISABLE HOST AUTHORIZATION COMPLETELY IN DEVELOPMENT
  config.hosts.clear
  
  # Explicitly allow HackerRank and common development hosts
  config.hosts << "vm-50010c11-7532-40c7-9fa2-e2166e365fd5-8000.us-vmprovider.projects.hrcdn.net"
  config.hosts << /.*\.hrcdn\.net/
  config.hosts << /.*\.projects\.hrcdn\.net/
  config.hosts << "localhost"
  config.hosts << "127.0.0.1"
  config.hosts << "0.0.0.0"

  # Use an evented file watcher to see changes in real-time.
  config.file_watcher = ActiveSupport::EventedFileUpdateChecker

  # Uncomment if you wish to allow Action Cable access from any origin.
  # config.action_cable.disable_request_forgery_protection = true
end 