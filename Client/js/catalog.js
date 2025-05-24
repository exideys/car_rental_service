document.addEventListener('DOMContentLoaded', () => {
  // 1) Подключаем CSS (если действительно нужен здесь — иначе можно оставить в <head>)
  const link = document.createElement('link');
  link.rel = 'stylesheet';
  link.href = 'card.css';
  document.head.appendChild(link);

  // 2) Получаем контейнер списка и рендерим автомобили при загрузке
  const list = document.querySelector('.car-list');
  fetch('/api/cars')
    .then(res => res.json())
    .then(cars => renderCars(cars, list));

  // 3) Обработчик для формы фильтра (на страницах без неё просто пропустит)
  const form = document.getElementById('filter-form');
  if (form) {
    form.addEventListener('submit', async e => {
      e.preventDefault();
      const formData = new FormData(form);
      const params = new URLSearchParams(formData);
      const resp = await fetch(`/api/cars?${params.toString()}`);
      const cars = await resp.json();
      renderCars(cars, list);
    });
  }

  // 4) Модалка «Замовити»
  const modal = document.getElementById('orderModal');
  if (modal) {
    const closeBtn = modal.querySelector('.modal-close');
    document.body.addEventListener('click', e => {
      if (e.target.matches('.order-button')) {
        modal.classList.add('show');
      } else if (e.target === closeBtn || e.target === modal) {
        modal.classList.remove('show');
      }
    });
  }
});

// Вынесем рендер в функцию, принимающую контейнер
function renderCars(cars, container) {
  container.innerHTML = '';
  cars.forEach(car => {
    const card = document.createElement('div');
    card.className = 'car-card';
    card.innerHTML = `
      <h3>${car.brand} ${car.model}</h3>
      <div class="car-card-content">
        <img src="${car.image_path}" alt="car image" />
        <div class="car-details">
          <p>Рік випуску: ${car.year_of_issue}</p>
          <p>Номер: ${car.plate_number}</p>
          <p>Статус: ${car.status}</p>
          <p>Класс: ${car.car_class}</p>
          <p>Ціна: ${car.daily_price}₴</p>
        </div>
      </div>
      <button class="order-button">Замовити</button>
    `;
    container.appendChild(card);
  });
}
