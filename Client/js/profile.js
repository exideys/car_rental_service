document.addEventListener('DOMContentLoaded', () => {
    fetch('/api/current_user')
            .then(res => {
              if (!res.ok) throw new Error('Unauthorized');
              return res.json();
            })
            .then(user => {
              renderProfile(user);
              fetchOrders(user.email);
            })
            .catch(() => {
              window.location.href = '/html/login.html';
            });

    document.getElementById('logout-btn')
            .addEventListener('click', () => {
              fetch('/logout', { method: 'POST' })
                      .then(() => window.location.href = '/html/login.html');
            });
  });

  function renderProfile(user) {
    document.querySelector('#profile-name').textContent =
            user.username ?? `${user.first_name} ${user.last_name}`;
    document.querySelector('#profile-email').textContent = user.email;
    document.querySelector('#profile-telephone').textContent = user.telephone || '—';
    document.querySelector('#profile-birth').textContent = user.birth_date;
    document.querySelector('#profile-created').textContent = user.created_at;
    document.querySelector('#profile-vip').textContent = user.is_vip ? 'Yes' : 'No';
  }

  function fetchOrders(email) {
    fetch(`/orders?email=${encodeURIComponent(email)}`, { method: 'GET' })
            .then(res => {
              if (!res.ok) throw new Error('Failed to load orders');
              return res.json();
            })
            .then(orders => {
              const carFetches = orders.map(order =>
                      fetch(`/car_id?car_id=${order.car_id}`)
                              .then(res => {
                                if (!res.ok) throw new Error('Car not found');
                                return res.json();
                              })
              );

              Promise.all(carFetches)
                      .then(cars => renderOrders(orders, cars))
                      .catch(err => {
                        console.error(err);
                        const tbody = document.querySelector('#orders-table tbody');
                        tbody.innerHTML = '<tr><td colspan="5">Error loading car details</td></tr>';
                      });
            })
            .catch(err => {
              console.error(err);
              const tbody = document.querySelector('#orders-table tbody');
              tbody.innerHTML = '<tr><td colspan="5">Error loading orders</td></tr>';
            });
  }

  function renderOrders(orders, cars) {
    const tbody = document.querySelector('#orders-table tbody');
    if (!orders.length) {
      tbody.innerHTML = '<tr><td colspan="5">You have no orders yet</td></tr>';
      return;
    }
    tbody.innerHTML = '';
    orders.forEach((order, idx) => {
      const car = cars[idx];
      const carName = `${car.brand} ${car.model}`; 
      const tr = document.createElement('tr');
      tr.innerHTML = `
          <td>${idx + 1}</td>
          <td>${carName}</td>
          <td>${order.start_date ?? order.StartDate}</td>
          <td>${order.end_date ?? order.EndDate}</td>
          <td>${order.status ?? order.status}</td>
        `;
      tbody.appendChild(tr);
    });
  }