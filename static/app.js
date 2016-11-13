angular.module('moolinet', ['ngResource'])
  .controller('ChallengeListController', function() {
    this.list = [
      {"title": "Hello world", "description": "bonjour le monde"},
      {"title": "Hello world 1", "description": "bonjour le monde"},
      {"title": "Hello world 2", "description": "bonjour le monde"},
      {"title": "Hello world 3", "description": "bonjour le monde"},
    ]
  })
;
