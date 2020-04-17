
$(document).ready(function () {

   calGridItemWidth()
});

function calGridItemWidth() {

    var screenWidth = $(document.body).width();
    var marginCenter = 20;
    var marginLeftAndRight = 30 * 2;

    var imgWidth = (screenWidth - marginCenter - marginLeftAndRight ) / 2;

    //宽高比例
    var scale = 3 / 4;
    var imgHeight = imgWidth * scale;
    $(".report-grid-img").css({ "width": imgWidth,  "height": imgHeight })
}
