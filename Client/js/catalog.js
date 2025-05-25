// catalog.js
document.addEventListener('DOMContentLoaded', () => {
  // --- 1) Подключаем CSS ---
  const link = document.createElement('link');
  link.rel = 'stylesheet';
  link.href = 'card.css';
  document.head.appendChild(link);

  // --- 2) Элементы страницы ---
  const list        = document.querySelector('.car-list');
  const filterForm  = document.getElementById('filter-form');
  const modal       = document.getElementById('orderModal');
  const closeBtn    = modal.querySelector('.modal-close');
  const summaryDiv  = modal.querySelector('#order-summary');
  const carIdInput  = modal.querySelector('input[name="car_id"]');
  const orderForm   = document.getElementById('order-form');
  const startInput  = orderForm.querySelector('input[name="start_date"]');
  const endInput    = orderForm.querySelector('input[name="end_date"]');
  const submitBtn   = orderForm.querySelector('button[type="submit"]');

  // --- 3) Загрузка и фильтрация ---
  async function loadCars(params = '') {
    const cars = await fetch(`/api/cars${params}`).then(r => r.json());
    renderCars(cars);
  }
  loadCars();
  filterForm.addEventListener('submit', e => {
    e.preventDefault();
    const qs = new URLSearchParams(new FormData(filterForm)).toString();
    loadCars(`?${qs}`);
  });

  // --- 4) Закрытие модалки ---
  closeBtn.addEventListener('click', () => modal.classList.remove('show'));
  modal.addEventListener('click', e => {
    if (e.target === modal) modal.classList.remove('show');
  });

  // --- 5) Валидация дат ---
  function validateDates() {
    if (startInput.value && endInput.value && startInput.value > endInput.value) {
      endInput.setCustomValidity('End Date не може бути раніше Start Date');
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

  // --- 6) Сабмит заказа ---
  orderForm.addEventListener('submit', async e => {
    e.preventDefault();
    const carId = carIdInput.value;

    validateDates();
    if (!orderForm.checkValidity()) return;

    // Проверяем авторизацию
    const userResp = await fetch('/api/current_user');
    if (!userResp.ok) {
      alert('Будь ласка, увійдіть перед оформленням замовлення.');
      return;
    }
    const user = await userResp.json();
    const fd = new FormData(orderForm);
    const dailyPrice = Number(fd.get('daily_price'));
    const payload = {
      email:      user.email,
      car_id:     Number(fd.get('car_id')),
      start_date: fd.get('start_date'),
      end_date:   fd.get('end_date'),
      daily_price: dailyPrice,
    };

    try {
      const resp = await fetch('/order', {
        method:  'POST',
        headers: { 'Content-Type': 'application/json' },
        body:    JSON.stringify(payload),
      });
      if (!resp.ok) {
        const err = await resp.json().catch(() => ({}));
        alert(err.error || 'Не вдалося створити замовлення.');
        return;
      }
      const order = await resp.json();
      alert(`Замовлення #${order.order_id} успішно створено.`);
      orderForm.reset();
      modal.classList.remove('show');
    } catch {
      alert('Помилка мережі. Спробуйте пізніше.');
    }
  });

  // --- 7) Рендер карточек и обработка клика ---
  function renderCars(cars) {
    list.innerHTML = '';
    cars.forEach(car => {
      const card = document.createElement('div');
      card.className = 'car-card';
      card.innerHTML = `
        <h3>${car.brand} ${car.model}</h3>
        <div class="car-card-content">
          <img src="${car.image_path}" alt="">
          <div class="car-details">
            <p>Рік: ${car.year_of_issue}</p>
            <p>Номер: ${car.plate_number}</p>
            <p>Статус: ${car.status}</p>
            <p>Клас: ${car.car_class}</p>
            <p>Ціна: ${car.daily_price}₴</p>
          </div>
        </div>
        <button class="order-button">Замовити</button>
      `;
      list.appendChild(card);
      const dailyPriceInput = orderForm.querySelector('#dailyPriceInput');
      const userIDInput = orderForm.querySelector('#clientIDInput');
      const btn = card.querySelector('.order-button');
      btn.addEventListener('click', async () => {
        // вот здесь используем корректное свойство
        const carId = car.car_id;  // <-- вместо car.id
        carIdInput.value = carId;

        // параллельно запрашиваем клиента и машину
        const [ userResp, carResp ] = await Promise.all([
          fetch('/api/current_user'),
          fetch(`/car_id?car_id=${carId}`)
        ]);
        const userData = userResp.ok ? await userResp.json() : {};
        const carData  = carResp.ok  ? await carResp.json()  : {};
        dailyPriceInput.value = carData.daily_price;
        userIDInput.value = userData.client_id;
        summaryDiv.innerHTML = `
          <h4>Підтвердження замовлення</h4>
          <p><b>Клієнт:</b> ${userData.first_name || '-'} ${userData.last_name || '-'} (${userData.email || '-'})</p>
          <p><b>Авто:</b> ${carData.brand || '-'} ${carData.model || '-'}</p>
          <p><b>Номер:</b> ${carData.plate_number || '-'}</p>
          <p><b>Ціна/день:</b> ${carData.daily_price != null ? carData.daily_price + '₴' : '-'}</p>
        `;
        modal.classList.add('show');
      });
    });
  }
});
