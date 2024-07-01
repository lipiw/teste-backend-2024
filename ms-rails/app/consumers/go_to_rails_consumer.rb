# frozen_string_literal: true
class GoToRailsConsumer < ApplicationConsumer
  def consume
    messages.each do |message|
      data = message.payload.with_indifferent_access
      Product.create(data)
      puts "Mensagem processada: #{data}"
    end
  end
end