document.addEventListener('DOMContentLoaded', () => {
  // --- 1) Подключаем CSS, если нужно ---
  const link = document.createElement('link');
  link.rel = 'stylesheet';
  link.href = 'card.css';
  document.head.appendChild(link);

  // --- 2) Рендер списка машин + фильтрация ---
  const list = document.querySelector('.car-list');
  const filterForm = document.getElementById('filter-form');

  async function loadCars(params = '') {
    const resp = await fetch(`/api/cars${params}`);
    const cars = await resp.json();
    renderCars(cars, list);
  }

  // Загрузка при старте
  loadCars();

  // По сабмиту формы фильтрации
  filterForm.addEventListener('submit', e => {
    e.preventDefault();
    const qs = new URLSearchParams(new FormData(filterForm)).toString();
    loadCars(`?${qs}`);
  });

  // --- 3) Открытие/закрытие модалки ---
  const modal = document.getElementById('orderModal');
  const closeBtn = modal.querySelector('.modal-close');

  document.body.addEventListener('click', e => {
    // если кликнули «Замовити» на карточке
    if (e.target.matches('.order-button')) {
      // перед открытием вставляем в скрытое поле ID машины
      const carId = e.target.closest('.car-card').dataset.carId;
      document.getElementById('order-car-id').value = carId;
      modal.classList.add('show');
    }
    // закрыть по крестику или клику на фон
    else if (e.target === closeBtn || e.target === modal) {
      modal.classList.remove('show');
    }
  });

  // --- 4) Логика формы заказа ---
  const orderForm  = document.getElementById('order-form');
  const startInput = orderForm.querySelector('input[name="start_date"]');
  const endInput   = orderForm.querySelector('input[name="end_date"]');
  const submitBtn  = document.getElementById('order-submit');

  function validateDates() {
    if (startInput.value && endInput.value && startInput.value > endInput.value) {
      endInput.setCustomValidity('End Date не может быть раньше Start Date');
      submitBtn.disabled = true;
    } else {
      endInput.setCustomValidity('');
      submitBtn.disabled = false;
    }
  }

  startInput.addEventListener('change', () => {
    endInput.min = startInput.value;
    validateDates();
  });
  endInput.addEventListener('change', () => {
    startInput.max = endInput.value;
    validateDates();
  });

  orderForm.addEventListener('submit', async e => {
    e.preventDefault();
    validateDates();
    if (!orderForm.checkValidity()) return;

    // 1) Проверяем, залогинен ли пользователь
    const userResp = await fetch('/api/current_user');
    if (!userResp.ok) {
      alert('Please sign in before placing an order.');
      return;
    }
    const user = await userResp.json();
    console.log('Current user payload:', user);
    console.log('Email field:', user.email);
    const fd = new FormData(orderForm);
    const payload = {
      email:    user.email,
     // car_id:     Number(fd.get('car_id')),
      car_id:     1,
      start_date: fd.get('start_date'),
      end_date:   fd.get('end_date')
    };

    // 3) Отправляем POST
    try {

      const resp = await fetch('/order', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json',
                    'Accept':       'application/json'
        },

        body: JSON.stringify(payload)
      });

      if (!resp.ok) {
        const err = await resp.json().catch(() => ({}));
        alert(err.error || 'Failed to create order.');
        return;
      }

      const order = await resp.json();
      alert(`Order #${order.order_id} has been created successfully.`);
      orderForm.reset();
      modal.classList.remove('show');
    } catch {
      alert('Network error. Please try again later.');
    }
  });
});

// Вынесли рендер в отдельную функцию
function renderCars(cars, container) {
  container.innerHTML = '';
  cars.forEach(car => {
    const card = document.createElement('div');
    card.className = 'car-card';
    card.dataset.carId = car.id;  // <-- здесь сохраняем ID
    card.innerHTML = `
      <h3>${car.brand} ${car.model}</h3>
      <div class="car-card-content">
        <img src="${car.image_path}" alt="car image" />
        <div class="car-details">
          <p>Рік випуску: ${car.year_of_issue}</p>
          <p>Номер: ${car.plate_number}</p>
          <p>Статус: ${car.status}</p>
          <p>Клас: ${car.car_class}</p>
          <p>Ціна: ${car.daily_price}₴</p>
        </div>
      </div>
      <button class="order-button">Замовити</button>
    `;
    container.appendChild(card);
  });
}
