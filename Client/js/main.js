window.addEventListener('load', () => {
    const car1 = document.querySelector('.hero-car');
    const car2 = document.querySelector('.hero-car-2');
  
    // Добавим небольшой таймер, чтобы выглядело как движение
    setTimeout(() => {
      car1.classList.add('animate');
    }, 300);
  
    setTimeout(() => {
      car2.classList.add('animate');
    }, 600);
  });
  