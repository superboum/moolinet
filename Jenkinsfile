node {
  stage('Configure') {
    deleteDir()
    def workspace = pwd()
    env.GOPATH="${workspace}"
    env.PATH=env.PATH+":"+env.GOPATH+"/bin"
    sh 'mkdir -p bin pkg src src/github.com/superboum/moolinet'
    sh 'env'
  }

  stage('Clone') {
    dir('src/github.com/superboum/moolinet') {
        checkout scm
        sh 'go get -d -v ./...'
    }

    // That's a horrible hack to choose Docker version to use...
    // We keep it as we don't expect to have other weird dependencies...
    dir('src/github.com/docker/docker') {
      sh 'git checkout 667315576fac663bd80bbada4364413692e57ac6 > /dev/null'
    }
  }

  stage('Test') {
    dir('src/github.com/superboum/moolinet') {
        sh 'go test ./...'
    }
  }

  stage('Lint') {
    dir('src/github.com/superboum/moolinet') {
      sh 'go get -u github.com/alecthomas/gometalinter'
      sh 'gometalinter --install'
      sh 'gometalinter -j 1 -t --deadline 100s --exclude "Errors unhandled." --disable gotype ./...'
    } 
  }
}

