angular.module('moolinet', ['ngResource', 'ngRoute'])

  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/challenge', {
        templateUrl: 'partials/challenge.html',
        controller: 'ChallengeListController'
      })
      .when('/hall-of-fame', {
      })
      .when('/configuration', {
      })
      .when('/login', {
      })
      .otherwise({
        redirectTo: '/challenge'
      })
  }])

  .controller('ChallengeListController', function() {
    this.list = [
      {"title": "Hello world", "description": "bonjour le monde"},
      {"title": "Hello world 1", "description": "bonjour le monde"},
      {"title": "Hello world 2", "description": "bonjour le monde"},
      {"title": "Hello world 3", "description": "bonjour le monde"},
    ]
  })
;
