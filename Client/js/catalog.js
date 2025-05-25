// catalog.js
document.addEventListener('DOMContentLoaded', () => {
  // --- 1) Подключаем CSS ---
  const link = document.createElement('link');
  link.rel = 'stylesheet';
  link.href = '../css/card.css';
  document.head.appendChild(link);

  // --- 2) Элементы страницы ---
  const list           = document.querySelector('.car-list');
  const filterForm     = document.getElementById('filter-form');
  const modal          = document.getElementById('orderModal');
  const closeBtn       = modal.querySelector('.modal-close');
  const summaryDiv     = modal.querySelector('#order-summary');
  const carIdInput     = modal.querySelector('input[name="car_id"]');
  const orderForm      = document.getElementById('order-form');
  const startInput     = orderForm.querySelector('input[name="start_date"]');
  const endInput       = orderForm.querySelector('input[name="end_date"]');
  const submitBtn      = orderForm.querySelector('button[type="submit"]');
  const dailyPriceInput= orderForm.querySelector('#dailyPriceInput');
  const clientIDInput  = orderForm.querySelector('#clientIDInput');

  // --- 2.1) Поле с итоговой суммой ---
  const totalDiv = document.createElement('div');
  totalDiv.id = 'total-sum-container';
  totalDiv.style.flex = '1 1 100%';
  totalDiv.innerHTML = `
    <label for="total-sum"><b>Всього:</b></label>
    <div id="total-sum">0₴</div>
  `;
  orderForm.insertBefore(totalDiv, submitBtn);

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

  // --- 5) Подсчёт и валидация дат ---
  function updateTotal() {
    if (startInput.value && endInput.value && dailyPriceInput.value) {
      const d1 = new Date(startInput.value);
      const d2 = new Date(endInput.value);
      const msPerDay = 1000 * 60 * 60 * 24;
      let days = Math.ceil((d2 - d1) / msPerDay);
      if (days < 1) days = 1;
      const total = days * Number(dailyPriceInput.value);
      document.getElementById('total-sum').textContent = `${total}₴`;
    } else {
      document.getElementById('total-sum').textContent = '0₴';
    }
  }

  function validateDates() {
    if (startInput.value && endInput.value && startInput.value > endInput.value) {
      endInput.setCustomValidity('End Date не може бути раніше Start Date');
      submitBtn.disabled = true;
    } else {
      endInput.setCustomValidity('');
      submitBtn.disabled = false;
    }
    updateTotal();
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
    validateDates();
    if (!orderForm.checkValidity()) return;

    const userResp = await fetch('/api/current_user');
    if (!userResp.ok) {
      alert('Будь ласка, увійдіть перед оформленням замовлення.');
      return;
    }
    const user = await userResp.json();
    const fd = new FormData(orderForm);
    const payload = {
      email:       user.email,
      car_id:      Number(fd.get('car_id')),
      start_date:  fd.get('start_date'),
      end_date:    fd.get('end_date'),
      daily_price: Number(fd.get('daily_price')),
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
      document.getElementById('total-sum').textContent = '0₴';
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

      const btn = card.querySelector('.order-button');
      btn.addEventListener('click', async () => {
        const carId = car.car_id;
        carIdInput.value = carId;

        const [userResp, carResp] = await Promise.all([
          fetch('/api/current_user'),
          fetch(`/car_id?car_id=${carId}`)
        ]);
        const userData = userResp.ok ? await userResp.json() : {};
        const carData  = carResp.ok  ? await carResp.json()  : {};

        dailyPriceInput.value = carData.daily_price;
        clientIDInput.value    = userData.client_id;
        summaryDiv.innerHTML = `
          <h4>Підтвердження замовлення</h4>
          <p><b>Name:</b> ${userData.first_name || '-'}</p>
          <p><b>Surname:</b> ${userData.last_name  || '-'}</p>
          <p><b>Email:</b> ${userData.email         || '-'}</p>
          <p><b>Авто:</b> ${carData.brand || '-'} ${carData.model || '-'}</p>
          <p><b>Номер:</b> ${carData.plate_number || '-'}</p>
          <p><b>Ціна/день:</b> ${carData.daily_price != null ? carData.daily_price + '₴' : '-'}</p>
        `;
        document.getElementById('total-sum').textContent = '0₴';
        modal.classList.add('show');
      });
    });
  }
});
