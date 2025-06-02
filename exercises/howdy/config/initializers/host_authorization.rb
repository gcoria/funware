# Disable host authorization for development and testing environments
# This initializer runs after application configuration is loaded

Rails.application.configure do
  # Clear all host restrictions
  config.hosts.clear
  
  # Explicitly allow HackerRank hosts
  config.hosts << "vm-50010c11-7532-40c7-9fa2-e2166e365fd5-8000.us-vmprovider.projects.hrcdn.net"
  config.hosts << /.*\.hrcdn\.net/
  config.hosts << /.*\.projects\.hrcdn\.net/
  config.hosts << /.*\.us-vmprovider\.projects\.hrcdn\.net/
  config.hosts << "localhost"
  config.hosts << "127.0.0.1"
  config.hosts << "0.0.0.0"
end

# Alternative approach: Disable the middleware entirely
if Rails.env.development? || Rails.env.test?
  Rails.application.config.middleware.delete ActionDispatch::HostAuthorization
end 