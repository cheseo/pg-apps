pipeline {
	agent any
	stages {
		stage ('all') {
		environment {
			AWS_ACCESS_KEY_ID = credentials('aws-access-key')
			AWS_SECRET_ACCESS_KEY = credentials('aws-secret-access-key')
			AWS_DEFAULT_REGION = "ap-south-1"
		}
		steps {
			dir("/home/public/cloud-scripts/week-7/jenkins-terra"){
			sh '''
			git pull
			terraform init
			terraform plan
			'''
			input message: "Continue?"
			sh '''
			terraform apply -auto-approve
			'''
			}
		}
		}
	}
}
