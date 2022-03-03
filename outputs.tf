# Output value definitions

output "function_name" {
  description = "Name of the lambda function"
  value       = aws_lambda_function.fight-irl.function_name
}

output "base_url" {
  description = "Base URL for API Gateway stage"
  value       = aws_apigatewayv2_stage.lambda.invoke_url
}
