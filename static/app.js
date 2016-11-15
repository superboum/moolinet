angular.module('moolinet', ['ngResource', 'ngRoute'])
  .config(function($resourceProvider) {
      $resourceProvider.defaults.stripTrailingSlashes = false;
  })

  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/challenge', {
        templateUrl: 'partials/challenge.html',
        controller: 'ChallengeListController'
      })
      .when('/challenge/:slug', {
        templateUrl: 'partials/challenge_view.html',
        controller: 'ChallengeViewController'
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

  .factory('Challenge', ['$resource', function($resource) {
    var Challenge = $resource('/api/challenge/:slug', {slug:'@id'});
    var list = []

    return {
      load: function(cb) { if (list.length == 0) { list = Challenge.query(cb) } else { cb(); } },
      getList: function() { return list; },
      getChallenge: function(slug) { return list.find(function(elem) { return elem.Slug == slug }); }
    }
  }])

  .factory('Job', ['$resource', function($resource) {
    var Job = $resource('/api/job/:uuid', {uuid:'@id'});
    return {
      submit: function(slug, vars, cb) { Job.save({Slug: slug, Vars: vars}, cb); },
      get: function(UUID, cb) { return Job.get({uuid: UUID}, cb); }
    };
  }])

  .filter('to_trusted', ['$sce', function($sce){
    return function(text) {
      return $sce.trustAsHtml(text);
    };
  }])

  .controller('ChallengeListController', ['Challenge', function(Challenge) {
    Challenge.load(function() { this.list = Challenge.getList(); }.bind(this));
  }])

  .controller('ChallengeViewController', ['Challenge', 'Job', '$routeParams', "$scope", "$interval", function(Challenge, Job, $routeParams, $scope, $interval) {
    $scope.status_list = ["WAITING IN QUEUE", "PROVISIONNING", "IN PROGRESS", "SUCCESS", "FAILED"];
    $scope.job = null;

    Challenge.load(function() {
      this.selected = Challenge.getChallenge($routeParams.slug);
      if (this.selected) { this.selected.Body = marked(this.selected.Body); }
    }.bind(this));

    $scope.submitAnswer = function() {
      Job.submit($routeParams.slug, {"[CODE]": $scope.code}, function(res) {
        $scope.job = res;
        finished = $interval(function() {
          Job.get(res.UUID, function(job) {
            $scope.job = job;
            if ($scope.job.Status >= 3) {
              $interval.cancel(finished);
            }
          });
        }, 1000);
      });
    };
  }])

  // See http://stackoverflow.com/a/25655639/6518667
  .directive('ngAllowTab', function () {
    return function (scope, element, attrs) {
      element.bind('keydown', function (event) {
        if (event.which == 9) {
          event.preventDefault();
          var start = this.selectionStart;
          var end = this.selectionEnd;
          element.val(element.val().substring(0, start) + '\t' + element.val().substring(end));
          this.selectionStart = this.selectionEnd = start + 1;
          element.triggerHandler('change');
        }
      });
    };
  });
