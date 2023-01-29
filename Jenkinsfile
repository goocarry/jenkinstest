/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.19.1-alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'apk add --no-cache make'
                sh 'make pgTestSetup'
                sh 'make pgTestStart'
            }
        }
        stage('test') {
            steps {
                sh 'make test'
            }
        }
    }
}