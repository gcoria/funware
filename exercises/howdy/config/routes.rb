Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  
  # Location API endpoint
  get '/whereami', to: 'location#whereami'
  
  # Report endpoints (based on existing controller)
  resources :reports, only: [:index, :show, :create]
end 