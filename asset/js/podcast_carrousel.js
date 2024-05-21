// Init le Carrousel avec Swiper
const swiper = new Swiper('.swiper', {
    slidesPerView: 3,
    spaceBetweenSlides: 0,
    direction: 'horizontal',
    loop: true,
    centeredSlides: true,
    // Scrollbar
    scrollbar: {
      el: '.swiper-scrollbar',
      hide: true,
    },
    // Flèches suivantes et précédentes
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
    breakpoints: {
      300: {
        slidesPerView: 1,
      },
      600: {
        slidesPerView: 2,
    },
    800: {
      slidesPerView: 3,
  },
}
});
