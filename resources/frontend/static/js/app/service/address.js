'use strict';

(function () {
	var Address = function($resource) {
		return $resource('/user/me/addresses/:id');
	};

	angular.module('App').factory(
		'Address', [
			'$resource',
			Address
		]
	);
}());