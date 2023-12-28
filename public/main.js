function openModal() {
  var modal = document.getElementById('loginModal');
  var overlay = document.querySelector('.modal-overlay');
  modal.classList.add('active');
  overlay.style.display = 'flex';
  document.body.style.overflow = 'hidden';
}

function closeModal() {
  var modal = document.getElementById('loginModal');
  var overlay = document.querySelector('.modal-overlay');
  modal.classList.remove('active');
  overlay.style.display = 'none';
  document.body.style.overflow = 'auto';
}

// Menutup modal jika klik di luar modal
window.onclick = function(event) {
  var modal = document.getElementById('loginModal');
  if (event.target == modal) {
    closeModal();
  }
};