resource "aws_ssm_parameter" "domain" {
  name = "domain"
  type = "String"
  value = "${var.domain}"
}
