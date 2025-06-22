pipeline{
    agent any

    environment {
        IMAGE_NAME= 'go-api-local'
        IMAGE_TAG = 'latest'
        CONTAINER_NAME = 'go-pos-toko-container'
    }

    stages{
        stage('Checkout'){
            steps {
                  git branch: 'main', url: 'https://github.com/fn199516/pos-toko.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME:$IMAGE_TAG .'
            }
        }

        stage('Stop Existing Container'){
            steps{
                sh '''
                docker stop $CONTAINER_NAME || true
                docker rm $CONTAINER_NAME || true
                '''
            }
        }

        stage('Run Container'){
            steps{
                sh 'docker run -d --name $CONTAINER_NAME -p 1010:1010 $IMAGE_NAME:$IMAGE_TAG'
            }
        }

    }


}