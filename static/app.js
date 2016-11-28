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
        templateUrl: 'partials/ranking.html',
        controller: 'RankingController'
      })
      .when('/login', {
        templateUrl: 'partials/login.html',
        controller: 'LoginController'
      })
      .otherwise({
        redirectTo: '/challenge'
      })
  }])

  .factory('CheckError', ['$location', function($location) {
    return function(error) {
      if (error.status == 401) {
        $location.path('/login');
      }
    }
  }])

  .factory('Challenge', ['$resource', 'CheckError', function($resource, CheckError) {
    var Challenge = $resource('/api/challenge/:slug', {slug:'@id'});
    var list = []

    return {
      load: function(cb) { if (list.length == 0) { list = Challenge.query(cb, CheckError) } else { cb(); } },
      getList: function() { return list; },
      getChallenge: function(slug) { return list.find(function(elem) { return elem.Slug == slug }); }
    }
  }])

  .factory('Job', ['$resource', 'CheckError', function($resource, CheckError) {
    var Job = $resource('/api/job/:uuid', {uuid:'@id'});
    return {
      submit: function(slug, vars, cb) { Job.save({Slug: slug, Vars: vars}, cb, CheckError); },
      get: function(UUID, cb) { return Job.get({uuid: UUID}, cb, CheckError); }
    };
  }])

  .factory('Me', [function() {
    var me = {Username: null, Created: null};
    return {
      set: function(u) { me.Username = u.Username; me.Created = u.Created; },
      get: function() { return me; },
      reset: function() { me.Username = null; me.Created = null;  }
    };
  }])

  .factory('Ranking', ['$http', 'CheckError', function($http, CheckError) {
    var raw_ranking = function(cb) {
      $http.get('/api/job/ranking').then(function(res) { cb(res.data); }, CheckError);
    };

    var computed_ranking = function(cb) {
      raw_ranking(function(raw) {
        res = [];
        for (var key in raw) {
          res.push({username: key, score: raw[key].length, last_submission: raw[key].reduce(function(acc, el) { return acc < new Date(el.Created) ? new Date(el.Created) : acc }, new Date("1970-01-01"))});
        }
        cb(res);
      });
    };

    return {
      raw: raw_ranking,
      computed: computed_ranking
    };
  }])

  .filter('to_trusted', ['$sce', function($sce){
    return function(text) {
      return $sce.trustAsHtml(text);
    };
  }])

  .controller('LoginController', ['$scope', '$http', '$location', 'Me', function($scope, $http, $location, Me) {
    $scope.connect = function() {
      send = {Username: $scope.login_username, Password: $scope.login_password}
      $http.post('/api/auth/login', send).then(function(response) {
        Me.set(response.data);
        $location.path('/');
      }, function(error) {
        alert(JSON.stringify(error.data.ErrDescription));
      });
    };

    $scope.register = function() {
      if ($scope.registration_password_1 != $scope.registration_password_2) {
        alert("passwords don't match");
        return
      }

      send = {Username: $scope.registration_username, Password: $scope.registration_password_1}
      $http.post('/api/auth/register', send).then(function(response) {
        Me.set(response.data);
        $location.path('/');
      }, function(error) {
        alert(JSON.stringify(error.data.ErrDescription))
      });
    };
  }])

  .controller('HeadController', ['Me', 'CheckError', '$http', '$scope', '$location', function(Me, CheckError, $http, $scope, $location) {
    $http.get('/api/auth/me').then(function(res) { Me.set(res.data); }, CheckError);
    $scope.me = Me.get();

    $scope.logout = function() {
      $http.get('/api/auth/logout').then(function(res) {
        Me.reset();
        $location.path('/login');
      });
    };
  }])

  .controller('RankingController', ['Ranking', '$scope', '$interval', function(Ranking, $scope, $interval) {
    f = function() {
      Ranking.computed(function(res) {
        $scope.ranking = res;
      });
    };
    finished = $interval(f, 10000);
    f();

    $scope.$on('$destroy',function(){
      $interval.cancel(finished);
    });
  }])

  .controller('ChallengeListController', ['Challenge', '$scope', function(Challenge, $scope) {
    $scope.chalList = []
    Challenge.load(function() { $scope.chalList = Challenge.getList(); }.bind(this));
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
