document.addEventListener("DOMContentLoaded", () => {
  const startInput = document.getElementById("start-date");
  const endInput = document.getElementById("end-date");
  const totalSum = document.getElementById("total-sum");
  const submitBtn = document.getElementById("submit-order");
  const dailyPriceInput = document.getElementById("daily-price");
  const carIdInput = document.getElementById("car-id");
  const orderModal = document.getElementById("order-modal");
  const closeModal = document.querySelector(".close-button");
  const orderForm = document.getElementById("order-form");

  if (startInput && endInput && totalSum && submitBtn && dailyPriceInput && carIdInput && orderModal && closeModal && orderForm) {
    startInput.addEventListener("change", validateDates);
    endInput.addEventListener("change", validateDates);
    startInput.addEventListener("change", updateTotal);
    endInput.addEventListener("change", updateTotal);

    closeModal.addEventListener("click", () => {
      orderModal.classList.remove("show");
    });

    window.addEventListener("click", (event) => {
      if (event.target === orderModal) {
        orderModal.classList.remove("show");
      }
    });

    orderForm.addEventListener("submit", async (e) => {
      e.preventDefault();

      const formData = new FormData(orderForm);
      const orderData = Object.fromEntries(formData.entries());

      try {
        const response = await fetch("/api/order", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(orderData),
        });

        if (response.ok) {
          alert("Order placed successfully!");
          orderModal.classList.remove("show");
          window.location.reload();
        } else {
          const errorData = await response.json();
          alert(`Error: ${errorData.message}`);
        }
      } catch (error) {
        console.error("Error placing order:", error);
        alert("An error occurred while placing the order.");
      }
    });
  }

  const carCards = document.querySelectorAll(".car-card");
  if (carCards) {
    carCards.forEach((card) => {
      card.addEventListener("click", async () => {
        const carId = card.dataset.carId;
        try {
          const response = await fetch(`/api/car/${carId}`);
          if (!response.ok) {
            throw new Error("Failed to fetch car details");
          }
          const carData = await response.json();
          openOrderModal(carData);
        } catch (error) {
          console.error("Error fetching car details:", error);
          alert("Failed to load car details.");
        }
      });
    });
  }

  function openOrderModal(carData) {
    const modal = document.getElementById("order-modal");
    if (!modal) return;

    document.getElementById("modal-car-image").src = carData.image_url || "/assets/default-car.png";
    document.getElementById("modal-car-name").textContent = `${carData.brand} ${carData.model}`;
    document.getElementById("modal-car-year").textContent = carData.year;
    document.getElementById("modal-car-price").textContent = `${carData.daily_price}₴/day`;
    document.getElementById("daily-price").value = carData.daily_price;
    document.getElementById("car-id").value = carData.id;

    // Reset dates and total
    startInput.value = "";
    endInput.value = "";
    totalSum.textContent = "0₴";

    modal.classList.add("show");
  }

  const filterForm = document.getElementById("filter-form");
  if (filterForm) {
    filterForm.addEventListener("submit", async (e) => {
      e.preventDefault();
      const formData = new FormData(filterForm);
      const queryParams = new URLSearchParams();
      for (const [key, value] of formData.entries()) {
        if (value) {
          queryParams.append(key, value);
        }
      }
      window.location.href = `/html/catalog.html?${queryParams.toString()}`;
    });
  }

  const clearFilterBtn = document.getElementById("clear-filter-btn");
  if (clearFilterBtn) {
    clearFilterBtn.addEventListener("click", () => {
      window.location.href = "/html/catalog.html";
    });
  }

  const searchInput = document.getElementById("search-input");
  if (searchInput) {
    searchInput.addEventListener("keypress", (e) => {
      if (e.key === "Enter") {
        e.preventDefault();
        const query = searchInput.value;
        window.location.href = `/html/catalog.html?search=${encodeURIComponent(query)}`;
      }
    });
  }

  const sortSelect = document.getElementById("sort-select");
  if (sortSelect) {
    sortSelect.addEventListener("change", (e) => {
      const sortValue = e.target.value;
      const url = new URL(window.location.href);
      url.searchParams.set("sort", sortValue);
      window.location.href = url.toString();
    });
  }

  const paginationLinks = document.querySelectorAll(".pagination-link");
  if (paginationLinks) {
    paginationLinks.forEach((link) => {
      link.addEventListener("click", (e) => {
        e.preventDefault();
        const page = link.dataset.page;
        const url = new URL(window.location.href);
        url.searchParams.set("page", page);
        window.location.href = url.toString();
      });
    });
  }

  const orderButtons = document.querySelectorAll(".order-button");
  if (orderButtons) {
    orderButtons.forEach((button) => {
      button.addEventListener("click", async (e) => {
        e.stopPropagation();
        const carId = button.dataset.carId;
        const modal = document.getElementById("order-modal");
        const carImage = modal.querySelector("#modal-car-image");
        const carName = modal.querySelector("#modal-car-name");
        const carYear = modal.querySelector("#modal-car-year");
        const carPrice = modal.querySelector("#modal-car-price");
        const dailyPriceInput = modal.querySelector("#daily-price");
        const carIdInput = modal.querySelector("#car-id");

        try {
          const response = await fetch(`/api/car/${carId}`);
          if (!response.ok) {
            throw new Error("Failed to fetch car details");
          }
          const carData = await response.json();

          carImage.src = carData.image_url || "/assets/default-car.png";
          carName.textContent = `${carData.brand} ${carData.model}`;
          carYear.textContent = carData.year;
          carPrice.textContent = `${carData.daily_price}₴/day`;
          dailyPriceInput.value = carData.daily_price;
          carIdInput.value = carData.id;

          // Reset dates and total
          startInput.value = "";
          endInput.value = "";
          document.getElementById("total-sum").textContent = "0₴";
          modal.classList.add("show");
        } catch (error) {
          console.error("Error fetching car details:", error);
          alert("Failed to load car details.");
        }
      });
    });
  }
});

function updateTotal() {
  const startInput = document.getElementById("start-date");
  const endInput = document.getElementById("end-date");
  const dailyPriceInput = document.getElementById("daily-price");
  const totalSum = document.getElementById("total-sum");

  if (startInput && endInput && dailyPriceInput && totalSum && startInput.value && endInput.value && dailyPriceInput.value) {
    const d1 = new Date(startInput.value);
    const d2 = new Date(endInput.value);
    const msPerDay = 1000 * 60 * 60 * 24;
    let days = Math.ceil((d2 - d1) / msPerDay);
    if (days < 1) days = 1;
    const total = days * Number(dailyPriceInput.value);
    totalSum.textContent = `${total}₴`;
  } else if (totalSum) {
    totalSum.textContent = "0₴";
  }
}

function validateDates() {
  const startInput = document.getElementById("start-date");
  const endInput = document.getElementById("end-date");
  const submitBtn = document.getElementById("submit-order");

  if (startInput && endInput && submitBtn) {
    const startDate = new Date(startInput.value);
    const endDate = new Date(endInput.value);

    if (startDate && endDate && startDate >= endDate) {
      alert("End date must be after start date.");
      submitBtn.disabled = true;
    } else {
      submitBtn.disabled = false;
    }
  }
}

module.exports = { updateTotal, validateDates };