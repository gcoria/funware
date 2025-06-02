class ApplicationController
  # Basic Rails API controller
  # Add any application-wide logic here
  
  # Skip CSRF token verification for API requests
  # protect_from_forgery with: :null_session
  
  private
  
  def json_response(object, status = :ok)
    render json: object, status: status
  end
end 