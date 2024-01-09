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
    document.getElementById("mySidenav").style.width = "250px";
    document.getElementById("main").style.marginLeft = "250px";
    document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
  }
  
  /* Set the width of the side navigation to 0 and the left margin of the page content to 0, and the background color of body to white */
  function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
    document.getElementById("main").style.marginLeft = "0";
    document.body.style.backgroundColor = "white";
  } 