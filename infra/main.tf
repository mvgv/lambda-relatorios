provider "aws" {
  region = "us-east-1" # Substitua pela sua região
}

variable "TF_LAMBDA_ZIP_PATH" {
  type = string
}

resource "aws_lambda_function" "lambda-relatorios" {
  function_name = "lambda-relatorios"
  role         = aws_iam_role.lambda-relatorios.arn
  handler      = "main"
  runtime      = "provided.al2023"

  filename     = var.TF_LAMBDA_ZIP_PATH # Recupera o zip da lambda disponibilizado pela esteira

  environment {
    variables = {
      EXAMPLE_ENV_VAR = "example"
    }
  }
}

resource "aws_iam_role_policy" "lambda_exec_policy" {
  name = "crud-api-exec-role-policy"
  role = aws_iam_role.lambda-relatorios.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "dynamodb:*",
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Action": "sns:*",
            "Effect": "Allow",
            "Resource": "*"
        }     
      ]  
}  
EOF
}

resource "aws_iam_role" "lambda-relatorios" {
  name = "lambda-relatorios"
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda-relatorios" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda-relatorios.name
}

resource "aws_lambda_permission" "with_sns" {
  statement_id  = "AllowExecutionFromSNS"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda-relatorios.function_name
  principal     = "sns.amazonaws.com"
  source_arn    = "arn:aws:sns:us-east-1:101478099523:solicitar-relatorio"
}