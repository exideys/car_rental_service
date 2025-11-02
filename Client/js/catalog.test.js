document.body.innerHTML = `
  <div id='orderModal'>
    <div class='modal-close'></div>
    <div id='order-summary'></div>
    <form id='order-form'>
      <input name='car_id' />
      <input name='start_date' type='date' />
      <input name='end_date' type='date' />
      <input id='dailyPriceInput' value='100' />
      <input id='clientIDInput' />
      <div id='total-sum'>0₴</div>
      <button type='submit'></button>
    </form>
  </div>
`;

const { updateTotal, validateDates } = require('./catalog.js');

describe('catalog.js', () => {
  let startInput, endInput, totalSum, submitBtn, dailyPriceInput;

  beforeEach(() => {
    // Mock window.alert to prevent JSDOM errors
    jest.spyOn(window, 'alert').mockImplementation(() => {});

    document.body.innerHTML = `
      <input id="start-date" type="date">
      <input id="end-date" type="date">
      <input id="daily-price" value="100">
      <span id="total-sum">0₴</span>
      <button id="submit-order">Submit</button>
    `;
    
    startInput = document.getElementById('start-date');
    endInput = document.getElementById('end-date');
    dailyPriceInput = document.getElementById('daily-price');
    totalSum = document.getElementById('total-sum');
    submitBtn = document.getElementById('submit-order');
    
    // Reset inputs
    startInput.value = '';
    endInput.value = '';
    dailyPriceInput.value = '100';
    totalSum.textContent = '0₴';
    endInput.setCustomValidity('');
    submitBtn.disabled = false;
  });

  describe('updateTotal', () => {
    it('should calculate the total price correctly for a single day', () => {
      startInput.value = '2025-11-01';
      endInput.value = '2025-11-01';
      updateTotal();
      expect(totalSum.textContent).toBe('100₴');
    });

    it('should calculate the total price correctly for multiple days', () => {
      startInput.value = '2025-11-01';
      endInput.value = '2025-11-05';
      updateTotal();
      expect(totalSum.textContent).toBe('400₴');
    });

    it('should show 0 if dates are missing', () => {
      startInput.value = '2025-11-01';
      updateTotal();
      expect(totalSum.textContent).toBe('0₴');
    });
  });

  describe('validateDates', () => {
    it('should disable submit if end date is before start date', () => {
      startInput.value = '2025-11-05';
      endInput.value = '2025-11-01';
      validateDates();
      expect(window.alert).toHaveBeenCalledWith('End date must be after start date.');
      expect(submitBtn.disabled).toBe(true);
    });

    it('should enable submit if dates are valid', () => {
      startInput.value = '2025-11-01';
      endInput.value = '2025-11-05';
      validateDates();
      expect(endInput.validationMessage).toBe('');
      expect(submitBtn.disabled).toBe(false);
    });
  });
});
