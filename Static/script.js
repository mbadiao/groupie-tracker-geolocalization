const initSlider = () => {
    const slideButtons = document.querySelectorAll(".slider-wrapper .slide-button");
    const imageList = document.querySelector(".slider-wrapper .image-list");
    const SliderScrollbar = document.querySelector(".container .slider-scrollbar");
    const scrollbarThumb = SliderScrollbar.querySelector(".scrollbar-thumb")
    const maxScrollLeft = imageList.scrollWidth - imageList.clientWidth;

    scrollbarThumb.addEventListener("mousedown", (e) => {
        const startX = e.clientX;
        const thumbPosition = scrollbarThumb.offsetLeft;

        const handleMouseMove = (e) => {
            const deltaX = e.clientX - startX
            const newThumbPosition = thumbPosition + deltaX
            const maxThumbPosition = SliderScrollbar.getBoundingClientRect().width - scrollbarThumb.offsetWidth;

            const boundedPosition = Math.max(0, Math.min(maxThumbPosition, newThumbPosition))
            const scrollPosition = (boundedPosition / maxThumbPosition) * maxScrollLeft;

            scrollbarThumb.style.left = `${boundedPosition}px`
            imageList.scrollLeft = scrollPosition
        }

        const handleMouseUp = () => {
            document.removeEventListener("mousemove", handleMouseMove)
            document.removeEventListener("mouseup", handleMouseUp)
        }

        document.addEventListener("mousemove", handleMouseMove)
        document.addEventListener("mouseup", handleMouseUp)
    })


    slideButtons.forEach(button => {
        button.addEventListener("click", () =>{
            const direction = button.id === "prev-slide" ? -1 : 1;
            const scrollAmount = imageList.clientWidth * direction;
            imageList.scrollBy({ left: scrollAmount, behavior: "smooth" })
        })
    })

    const handleSlideButtons = () => {
        slideButtons[0].style.display = imageList.scrollLeft <= 0 ? "none" : "block";
        slideButtons[1].style.display = imageList.scrollLeft >= maxScrollLeft ? "none" : "block";
    }

    const updateScrollThumbPosition = () => {
        const scrollPosition = imageList.scrollLeft
        const thumbPosition = (scrollPosition / maxScrollLeft) * (SliderScrollbar.clientWidth - scrollbarThumb.offsetWidth)
        scrollbarThumb.style.left = `${thumbPosition}px`
   
    }

    imageList.addEventListener("scroll", () => {
        handleSlideButtons();
        updateScrollThumbPosition();
    })
}


window.addEventListener("load", initSlider)


/* Set the width of the side navigation to 250px and the left margin of the page content to 250px and add a black background color to body */
function openNav() {
    document.getElementById("mySidenav").style.width = "300px";
    document.getElementById("main").style.marginLeft = "250px";
    document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
  }
  
  /* Set the width of the side navigation to 0 and the left margin of the page content to 0, and the background color of body to white */
  function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
    document.getElementById("main").style.marginLeft = "0";
    document.body.style.backgroundColor = "white";
}


let rangeMin = 1963;
const range = document.querySelector(".range-selected");
const rangeInput = document.querySelectorAll(".range-input input");
const rangePrice = document.querySelectorAll(".range-price input");
const NumberOfArtist = document.querySelectorAll(".image-items")
console.log(NumberOfArtist)

rangeInput.forEach((input) => {
    input.addEventListener("input", (e) => {
        let minRange = parseInt(rangeInput[0].value);
        let maxRange = parseInt(rangeInput[1].value);
        if (e.target.className === "min") {
            if (minRange > maxRange) {
                rangeInput[1].value = minRange;
            }
        } else {
            if (maxRange < minRange) {
                rangeInput[0].value = maxRange;
            }
        }

        rangePrice[0].value = rangeInput[0].value;
        rangePrice[1].value = rangeInput[1].value;
        range.style.left = ((minRange - 1963) / (2018 - 1963)) * 100 + "%";
        range.style.right = 100 - ((maxRange - 1963) / (2018 - 1963)) * 100 + "%";
    });
});

rangePrice.forEach((input) => {
    input.addEventListener("input", (e) => {
        let minPrice = rangePrice[0].value;
        let maxPrice = rangePrice[1].value;
        if (e.target.className === "min" && minPrice > maxPrice) {
            rangeInput[1].value = minPrice;
        } else if (e.target.className === "max" && maxPrice < minPrice) {
            rangeInput[0].value = maxPrice;
        }
        range.style.left = ((rangeInput[0].value - 1963) / (2018 - 1963)) * 100 + "%"; 
        range.style.right = 100 - ((rangeInput[1].value - 1963) / (2018 - 1963)) * 100 + "%";
    });
});


let rangeMindate = 1958;
const rangedate = document.querySelector(".range-selecteds");
const rangeInputs = document.querySelectorAll(".range-inputs input");
const rangePrices = document.querySelectorAll(".range-prices input");

rangeInputs.forEach((input) => {
    input.addEventListener("input", (e) => {
        let minRangedate = parseInt(rangeInputs[0].value);
        let maxRangedate = parseInt(rangeInputs[1].value);
        if (e.target.className === "mindate") {
            if (minRangedate > maxRangedate) {
                rangeInputs[1].value = minRangedate;
            }
        } else {
            if (maxRangedate < minRangedate) {
                rangeInputs[0].value = maxRangedate;
            }
        }

        rangePrices[0].value = rangeInputs[0].value;
        rangePrices[1].value = rangeInputs[1].value;
        rangedate.style.left = ((minRangedate - 1958) / (2015 - 1958)) * 100 + "%"; 
        rangedate.style.right = 100 - ((maxRangedate - 1958) / (2015 - 1958)) * 100 + "%"; 
    });
});

rangePrices.forEach((input) => {
    input.addEventListener("input", (e) => {
        let minPricedate = rangePrices[0].value;
        let maxPricedate = rangePrices[1].value;
        if (e.target.className === "mindate" && minPricedate > maxPricedate) {
            rangeInputs[1].value = minPricedate;
        } else if (e.target.className === "maxdate" && maxPricedate < minPricedate) {
            rangeInputs[0].value = maxPricedate;
        }
        rangedate.style.left = ((rangeInputs[0].value - 1958) / (2015 - 1958)) * 100 + "%"; 
        rangedate.style.right = 100 - ((rangeInputs[1].value - 1958) / (2015 - 1958)) * 100 + "%"; 
    });
});
