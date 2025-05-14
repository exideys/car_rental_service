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
            `;
                list.append(card);
            });
        });
});