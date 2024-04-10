// Init le Carrousel avec Swiper
const swiper = new Swiper('.swiper', {
    slidesPerView: 3,
    spaceBetweenSlides: 0,
    direction: 'horizontal',
    loop: true,
    centeredSlides: true,
  
    // Pagination avec les points
    pagination: {
      el: '.swiper-pagination',
      clickable: true,
    },
  
    // Flèches suivantes et précédentes
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
});
