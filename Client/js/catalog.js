document.addEventListener('DOMContentLoaded', () => {
    fetch('/api/cars')
        .then(res => res.json())
        .then(cars => {
            const list = document.querySelector('.car-list');
            cars.forEach(car => {
                const card = document.createElement('div');
                card.className = 'car-card';
                card.innerHTML = `
              <h3>${car.brand} ${car.model}</h3>
              <p>Рік випуску: ${car.year_of_issue}</p>
              <p>Номер: ${car.plate_number}</p>
              <p>Статус: ${car.status}</p>
              <p>Ціна:${car.daily_price}<p>
            `;
                list.append(card);
            });
        });
});
function renderCars(cars) {
    const container = document.querySelector('.car-list');
    container.innerHTML = '';
    cars.forEach(car => {
        const carElement = document.createElement('div');
        carElement.className = 'car-item';
        carElement.innerHTML = `
      <h3>${car.brand} ${car.model}</h3>
              <p>Рік випуску: ${car.year_of_issue}</p>
              <p>Номер: ${car.plate_number}</p>
              <p>Статус: ${car.status}</p>
              <p>Ціна: ${car.daily_price}</p>
    `;
        container.appendChild(carElement);
    });
}
const form = document.getElementById('filter-form');
form.addEventListener('submit', async e => {
    e.preventDefault();
    const formData = new FormData(form);
    const params = new URLSearchParams(formData);
    const resp = await fetch(`/api/cars?${params.toString()}`);
    const cars = await resp.json();
    renderCars(cars);
});
