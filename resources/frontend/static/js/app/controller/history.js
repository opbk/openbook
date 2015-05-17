'use strict';

(function () {
	var HistoryController = function($scope, $http, toaster) {
		$scope.order_hide = {}
		$scope.book_show = {}

		$scope.delete = function(id) {
			if(confirm('Вы дествительно хотите удалить заказ?')) {
				$http.delete('/order/' + id)
					.success(function(data, status, headers, config) {
						console.log("order deleted successefully")
						$scope.order_hide[id] = true;
						toaster.pop('success', '', 'Заказ успешно удален')
					})
					.error(function(data, status, headers, config) {
						toaster.pop('error', '', 'При удалении заказа возникла ошибка, попробуйте повторить попытку через несколько минут')
					})
			}
		}

		$scope.return = function(id) {
			if(confirm('Вы дествительно хотите вернуть книгу')) {
				$http.get('/order/' + id + '/return')
					.success(function(data, status, headers, config) {
						console.log("order returned successefully")
						toaster.pop('success', '', 'Спасибо. В ближашее время с вами свяжуться для уточнения условия возврата')
					})
					.error(function(data, status, headers, config) {
						toaster.pop('error', '', 'При попытке возврата произошла ошибка, попробуйте повторить попытку через несколько минут')
					})
			}
		}
	};

	angular.module('App').controller(
		'HistoryController', [
			'$scope',
			'$http',
			'toaster',
			HistoryController
		]
	);
}());