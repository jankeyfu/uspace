// api = "/holiday/get"
// this.$http.get(api).then((response) => {
//     console.log(response.data);//成功回调
// }, (response) => {
//     //失败回调
// })
// $.ajax({
//     url: "/holiday/get",
//     data: {
//         user_id: 2
//     },
//     success: function (result) {
//         console.log(result)
//         // $("#weather-temp").html("<strong>" + result + "</strong> degrees");
//     }
// });
// 
// var app = new Vue({
//     el: '#holiday',
//     data: {
//         message: ''
//     },
//     created: function () {
//         var url = "/holiday/get";
//         this.$http.json(url, { user_id: 2 }).then(function (data) {
//             var json = data.body;
//             this.message = eval("(" + json + ")");
//         }, function (response) {
//             console.info(response);
//         })
//     }
// });