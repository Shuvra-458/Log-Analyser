pipeline {
    agent any

    environment {
        APP_NAME = "golog"
        DOCKER_IMAGE = "golog:latest"
        DOCKERHUB_REPO = "shuvra458/golog"
        GOTOOLCHAIN = "auto"
    }

    stages {

        stage('Checkout') {
            steps {
                echo "Checking out source code..."
                checkout scm
            }
        }

        stage('Go Build') {
            steps {
                echo "Building Go binary..."
                sh 'go version'
                sh 'go mod tidy'
                sh 'go build -o ${APP_NAME} .'
            }
        }

        stage('Docker Build') {
            steps {
                echo "Building Docker Image..."
                sh 'docker build -t ${DOCKER_IMAGE} .'
            }
        }

        stage('Docker Push') {
            steps {
                echo "Pushing image to DockerHub..."
                withCredentials([usernamePassword(credentialsId: 'docker-hub-cred', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh """
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker tag ${DOCKER_IMAGE} ${DOCKERHUB_REPO}:latest
                        docker push ${DOCKERHUB_REPO}:latest
                    """
                }
            }
        }
    }
}
