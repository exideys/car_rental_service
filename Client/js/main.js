window.addEventListener('load', () => {
    const car1 = document.querySelector('.hero-car');
    const car2 = document.querySelector('.hero-car-2');

    setTimeout(() => {
      car1.classList.add('animate');
    }, 200);
  
    setTimeout(() => {
      car2.classList.add('animate');
    }, 200);
  });
  