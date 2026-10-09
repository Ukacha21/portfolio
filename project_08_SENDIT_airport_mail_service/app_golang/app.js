/* ============================================================
   SENDIT — Airport Item Mail-Home Kiosk
   Dashboard App Logic v3.0
   - REST API backend (SQLite via Express)
   - Dual Mode: User (Kiosk) / Admin
   - Mobile QR form relay
   ============================================================ */

'use strict';

const API = `${window.location.origin}/api`;

// ===================== MODE MANAGEMENT =====================
const MODE = { USER: 'user', ADMIN: 'admin' };
let currentMode = MODE.USER;
let adminTimeout = null;
const ADMIN_TIMEOUT_MS = 5 * 60 * 1000; // 5 minutes

function switchToUserMode() {
  currentMode = MODE.USER;
  clearAdminTimeout();
  // Show user screen, hide admin panel
  document.getElementById('user-mode-screen').classList.add('active');
  document.getElementById('admin-panel').classList.remove('active');
  document.getElementById('admin-mode-badge').style.display = 'none';
  // Show user 3-dot menu, close dropdown
  const userDots = document.getElementById('user-menu-dots-wrap');
  if (userDots) userDots.style.display = '';
  document.getElementById('menu-dropdown').classList.remove('open');
  // Close any open QR prompt
  closeNewShipmentPromptDirect();
}

function enterAdminMode() {
  document.getElementById('menu-dropdown').classList.remove('open');
  openPinModal();
}

async function verifyAndSwitchAdmin(pin) {
  try {
    const resp = await fetch(`${API}/auth/verify-pin`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ pin }),
    });
    const data = await resp.json();
    if (data.ok) {
      currentMode = MODE.ADMIN;
      closePinModal();
      // Hide user mode screen & user 3-dot menu
      document.getElementById('user-mode-screen').classList.remove('active');
      const userDots = document.getElementById('user-menu-dots-wrap');
      if (userDots) userDots.style.display = 'none';
      // Show admin panel and badge
      document.getElementById('admin-panel').classList.add('active');
      document.getElementById('admin-mode-badge').style.display = 'flex';
      // Close any QR prompt
      closeNewShipmentPromptDirect();
      resetAdminTimeout();
      await refreshAll();
      // If a callback was queued (e.g. from user-mode new shipment flow), run it
      if (_afterAdminCallback) {
        const cb = _afterAdminCallback;
        _afterAdminCallback = null;
        showPage('dashboard');
        await cb();
      } else {
        showPage('dashboard');
      }
      showToast('Admin mode activated', 'success');
    } else {
      shakePinInput();
      showToast('Incorrect PIN', 'error');
    }
  } catch (e) {
    showToast('Server not reachable. Is node server.js running?', 'error');
  }
}

function resetAdminTimeout() {
  clearAdminTimeout();
  adminTimeout = setTimeout(() => {
    switchToUserMode();
    showToast('Admin session expired — returned to kiosk mode', 'info');
  }, ADMIN_TIMEOUT_MS);
}

function clearAdminTimeout() {
  if (adminTimeout) { clearTimeout(adminTimeout); adminTimeout = null; }
}

// ===================== PIN MODAL =====================
let pinValue = '';

function openPinModal() {
  pinValue = '';
  updatePinDots();
  document.getElementById('pin-error').textContent = '';
  document.getElementById('pin-overlay').classList.add('open');
  document.getElementById('pin-modal').classList.add('open');
}

function closePinModal() {
  pinValue = '';
  document.getElementById('pin-overlay').classList.remove('open');
  document.getElementById('pin-modal').classList.remove('open');
}

function pinKey(digit) {
  if (pinValue.length >= 8) return;
  pinValue += String(digit);
  updatePinDots();
  if (pinValue.length >= 4) {
    setTimeout(() => verifyAndSwitchAdmin(pinValue), 200);
  }
}

function pinBackspace() {
  pinValue = pinValue.slice(0, -1);
  updatePinDots();
}

function updatePinDots() {
  const dots = document.querySelectorAll('.pin-dot');
  dots.forEach((d, i) => {
    d.classList.toggle('filled', i < pinValue.length);
  });
}

function shakePinInput() {
  const dots = document.getElementById('pin-dots');
  pinValue = '';
  updatePinDots();
  dots.classList.add('shake');
  setTimeout(() => dots.classList.remove('shake'), 600);
}

// ===================== MENU DOTS =====================
function toggleMenu() {
  // Only toggle the user mode menu dropdown
  document.getElementById('menu-dropdown').classList.toggle('open');
}

document.addEventListener('click', (e) => {
  if (!e.target.closest('.menu-dots-btn') && !e.target.closest('.menu-dropdown')) {
    document.getElementById('menu-dropdown')?.classList.remove('open');
  }
});

// ===================== NEW SHIPMENT QR PROMPT =====================
let nspQrPollTimer  = null;
let nspQrCountTimer = null;
let nspQrCode       = null;

function openNewShipmentPrompt() {
  const overlay = document.getElementById('new-shipment-overlay');
  overlay.classList.add('open');
  // Show loading state
  document.getElementById('nsp-qr-loading').style.display = 'flex';
  document.getElementById('nsp-qr-content').style.display = 'none';
  // Generate QR session
  generateNspQR();
}

async function generateNspQR() {
  try {
    const resp = await fetch(`${API.replace('/api','')}/api/session/create`, { method: 'POST' });
    const data = await resp.json();
    if (!data.ok) throw new Error(data.error);
    nspQrCode = data.code;
    const mobileUrl  = data.mobile_url;
    const expiresAt  = new Date(data.expires_at);
    const qrImgUrl   = `https://api.qrserver.com/v1/create-qr-code/?size=260x260&data=${encodeURIComponent(mobileUrl)}&bgcolor=ffffff&color=1a1a2e&qzone=2`;
    const imgEl      = document.getElementById('nsp-qr-image');
    imgEl.src        = qrImgUrl;
    document.getElementById('nsp-qr-code-text').textContent = data.code;
    document.getElementById('nsp-qr-url-text').textContent  = mobileUrl.replace(/https?:\/\//, '');
    // Show content, hide loading
    document.getElementById('nsp-qr-loading').style.display = 'none';
    document.getElementById('nsp-qr-content').style.display  = 'flex';
    // Countdown
    nspQrCountTimer = setInterval(() => {
      const left = Math.max(0, Math.floor((expiresAt - Date.now()) / 1000));
      const m = Math.floor(left / 60), s = left % 60;
      const el = document.getElementById('nsp-qr-timer');
      if (el) el.textContent = `Expires in ${m}:${String(s).padStart(2,'0')}`;
      if (left === 0) { clearInterval(nspQrCountTimer); }
    }, 1000);
    // Poll for mobile submission
    nspQrPollTimer = setInterval(async () => {
      try {
        const pr = await fetch(`${API.replace('/api','')}/api/session/poll/${nspQrCode}`);
        const pd = await pr.json();
        if (pd.ok && pd.status === 'received') {
          clearInterval(nspQrPollTimer);
          clearInterval(nspQrCountTimer);
          // Data received — proceed to kiosk self-service confirmation screen
          _pendingMobileData = pd.data;
          closeNewShipmentPromptDirect();
          proceedToKioskConfirm();
        }
      } catch(e) {}
    }, 2000);
  } catch (e) {
    document.getElementById('nsp-qr-loading').innerHTML = `<span style="color:var(--warning);font-size:13px">⚠️ QR unavailable — please use staff form</span>`;
  }
}

let _pendingMobileData = null;

function closeNewShipmentPrompt(e) {
  if (e && e.target === document.getElementById('new-shipment-overlay')) {
    closeNewShipmentPromptDirect();
  }
}
function closeNewShipmentPromptDirect() {
  clearInterval(nspQrPollTimer);
  clearInterval(nspQrCountTimer);
  nspQrCode = null;
  const overlay = document.getElementById('new-shipment-overlay');
  if (overlay) overlay.classList.remove('open');
}

function startKioskSelfServiceFlow() {
  closeNewShipmentPromptDirect();
  resetForm();
  document.body.classList.add('user-kiosk-flow');
  currentMode = MODE.USER;
  document.getElementById('user-mode-screen').classList.remove('active');
  document.getElementById('admin-panel').classList.add('active');
  showPage('new-shipment');
}

function exitKioskSelfServiceFlow() {
  document.body.classList.remove('user-kiosk-flow');
  resetForm();
  switchToUserMode();
}

function proceedToKioskConfirm() {
  document.body.classList.add('user-kiosk-flow');
  currentMode = MODE.USER;
  document.getElementById('user-mode-screen').classList.remove('active');
  document.getElementById('admin-panel').classList.add('active');
  showPage('new-shipment');
  if (_pendingMobileData) {
    autoFillAllFromMobileData(_pendingMobileData);
    _pendingMobileData = null;
  }
  nextStep(4); // Go straight to Step 4: Payment & Review
}

function proceedToAdminShipment() {
  closeNewShipmentPromptDirect();
  // Open PIN modal to enter admin mode, then navigate to new shipment
  _afterAdminCallback = () => {
    showPage('new-shipment');
    if (_pendingMobileData) {
      autoFillFromMobileData(_pendingMobileData);
      nextStep(3); // Skip to package selection since address is pre-filled
      _pendingMobileData = null;
    }
  };
  openPinModal();
}

let _afterAdminCallback = null;

// ===================== API CALLS =====================
async function apiGet(path) {
  const resp = await fetch(`${API}${path}`);
  return resp.json();
}

async function apiPost(path, body) {
  const resp = await fetch(`${API}${path}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  return resp.json();
}

async function apiPut(path, body) {
  const resp = await fetch(`${API}${path}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  return resp.json();
}

async function apiDelete(path) {
  const resp = await fetch(`${API}${path}`, { method: 'DELETE' });
  return resp.json();
}

// ===================== DB STATUS =====================
async function checkServerStatus() {
  try {
    const r = await fetch(`${API}/stats`);
    if (r.ok) {
      setDBStatus(true);
      return true;
    }
  } catch (e) {}
  setDBStatus(false);
  return false;
}

function setDBStatus(ok) {
  const dot  = document.querySelector('.db-dot');
  const text = document.getElementById('db-status-text');
  if (ok) {
    dot.style.background = 'var(--success)';
    text.textContent = 'SQLite Connected';
  } else {
    dot.style.background = 'var(--danger)';
    text.textContent = 'Server Offline';
  }
}

// ===================== STATE =====================
const state = {
  currentStep: 1,
  currentPage: 'dashboard',
  selectedPackage: 'small',
  selectedCourier: 'dhl',
  insuranceAdded: false,
  selectedPayMethod: null,
  theme: 'dark-blue', textColor: 'white', accent: 'indigo', fontSize: 'medium',
  packageFilter: 'all',
};

const pricing = {
  packages:  { envelope: 1.50, small: 2.50, medium: 4.00 },
  couriers:  { dhl: { envelope:3.50,small:5.90,medium:9.50 }, dpd: { envelope:2.80,small:4.20,medium:7.20 }, rm: { envelope:2.20,small:3.50,medium:6.00 } },
  serviceFee: 2.50,
  insurance:  1.99,
};

// ===================== CURRENCY SYSTEM =====================
// Base prices are in GBP. Exchange rates are approximate display-only.
const CURRENCIES = {
  GBP: { symbol: '£',    name: 'British Pound',     rate: 1.000 },
  EUR: { symbol: '€',    name: 'Euro',               rate: 1.18  },
  USD: { symbol: '$',    name: 'US Dollar',           rate: 1.27  },
  CHF: { symbol: 'Fr',   name: 'Swiss Franc',         rate: 1.14  },
  AUD: { symbol: 'A$',   name: 'Australian Dollar',   rate: 1.95  },
  TRY: { symbol: '₺',    name: 'Turkish Lira',        rate: 40.5  },
  QAR: { symbol: 'QR',   name: 'Qatari Riyal',        rate: 4.62  },
  SAR: { symbol: 'SAR',  name: 'Saudi Riyal',         rate: 4.77  },
  BHD: { symbol: 'BD',   name: 'Bahraini Dinar',      rate: 0.478 },
  KWD: { symbol: 'KD',   name: 'Kuwaiti Dinar',       rate: 0.389 },
  EGP: { symbol: 'E£',   name: 'Egyptian Pound',      rate: 62.0  },
};

let activeCurrency = 'GBP';

function formatMoney(gbp) {
  const c = CURRENCIES[activeCurrency] || CURRENCIES.GBP;
  const converted = gbp * c.rate;
  const dec = (activeCurrency === 'KWD' || activeCurrency === 'BHD') ? 3 : 2;
  return `${c.symbol}${converted.toFixed(dec)}`;
}

function setCurrency(code) {
  if (!CURRENCIES[code]) return;
  activeCurrency = code;
  // Update any visible price labels right away
  updateCourierRates();
  updateTotals();
  // Highlight the active button
  document.querySelectorAll('.currency-btn').forEach(b => b.classList.toggle('active', b.dataset.currency === code));
  // Persist to DB
  fetch(`${API}/settings/currency`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ value: code }),
  }).catch(() => {});
  savePreferences();
}

const itemMeta = {
  liquid:      { icon:'💧', label:'Liquid / Aerosol / Gel',    eligible:true, hazmat:'Class 3 (minor)' },
  sharp:       { icon:'🔪', label:'Sharp Object',              eligible:true, hazmat:'None' },
  electronics: { icon:'🔋', label:'Electronics / Power Bank',  eligible:true, hazmat:'Class 9 (li-ion)' },
  lighter:     { icon:'🔥', label:'Lighter / Match',           eligible:true, hazmat:'Class 2.1' },
  cosmetics:   { icon:'💄', label:'Cosmetics / Toiletries',    eligible:true, hazmat:'None' },
  food:        { icon:'🍷', label:'Food / Beverage',           eligible:true, hazmat:'None' },
  tool:        { icon:'🔧', label:'Tool / Hardware',           eligible:true, hazmat:'None' },
  other:       { icon:'📦', label:'Other Item',                eligible:true, hazmat:'Unknown' },
};

const statusMeta = {
  pending: { label:'Pending / Available', color:'#f59e0b', bg:'rgba(245,158,11,0.12)',  icon:'⏳' },
  posted:  { label:'Posted / In Transit', color:'#6366f1', bg:'rgba(99,102,241,0.12)',  icon:'🚚' },
  mailed:  { label:'Mailed / Delivered',  color:'#10b981', bg:'rgba(16,185,129,0.12)', icon:'✅' },
  failed:  { label:'Failed / Returned',   color:'#ef4444', bg:'rgba(239,68,68,0.12)',  icon:'❌' },
};

// ===================== INIT =====================
document.addEventListener('DOMContentLoaded', async () => {
  loadPreferences();
  startClock();
  setInterval(updateQueueCount, 8000);
  setInterval(checkServerStatus, 30000);

  // Start in User Mode (ensures correct initial state)
  currentMode = MODE.USER;
  document.getElementById('user-mode-screen').classList.add('active');
  document.getElementById('admin-panel').classList.remove('active');
  document.getElementById('admin-mode-badge').style.display = 'none';
  const userDots = document.getElementById('user-menu-dots-wrap');
  if (userDots) userDots.style.display = '';

  // Try to connect to server
  const online = await checkServerStatus();
  if (online) {
    await refreshAll();
    // Load saved currency from DB
    try {
      const s = await apiGet('/settings');
      if (s.ok && s.data.currency && CURRENCIES[s.data.currency]) {
        activeCurrency = s.data.currency;
        document.querySelectorAll('.currency-btn').forEach(b => b.classList.toggle('active', b.dataset.currency === activeCurrency));
      }
    } catch(e) {}
    showToast('Database connected ✓', 'success');
  } else {
    showToast('⚠️ Server offline — run: node server.js', 'error');
  }

  // Courier radio change
  document.addEventListener('change', (e) => {
    if (e.target.name === 'courier') {
      state.selectedCourier = e.target.value;
      document.querySelectorAll('.courier-option').forEach(o => o.classList.remove('active'));
      const map = { dhl:'courier-dhl', dpd:'courier-dpd', rm:'courier-rm' };
      document.getElementById(map[e.target.value]).classList.add('active');
      updateTotals();
    }
  });

  // Reset admin timeout on user activity (within admin panel)
  document.addEventListener('mousemove', () => { if (currentMode === MODE.ADMIN) resetAdminTimeout(); });
  document.addEventListener('touchstart', () => { if (currentMode === MODE.ADMIN) resetAdminTimeout(); });
  document.addEventListener('keydown', () => { if (currentMode === MODE.ADMIN) resetAdminTimeout(); });
});

// ===================== REFRESH ALL =====================
async function refreshAll() {
  try {
    const result = await apiGet('/shipments');
    if (!result.ok) return;
    const all = result.data;
    updateStatusCounts(all);
    renderDashboardDonut(all);
    renderActivityList(all);
    renderCategoryBars(all);
    renderPackages(all);
  } catch (e) {
    console.error('refreshAll error:', e);
  }
}

// ===================== STATUS COUNTS =====================
function updateStatusCounts(all) {
  const counts = { pending:0, posted:0, mailed:0, failed:0 };
  all.forEach(s => { if (counts[s.status] !== undefined) counts[s.status]++; });
  const total = all.length;
  setText('belt-total',   total);
  setText('belt-pending', counts.pending);
  setText('belt-posted',  counts.posted);
  setText('belt-mailed',  counts.mailed);
  setText('belt-failed',  counts.failed);
  setText('fc-all',     total);
  setText('fc-pending', counts.pending);
  setText('fc-posted',  counts.posted);
  setText('fc-mailed',  counts.mailed);
  setText('fc-failed',  counts.failed);
}

// ===================== DASHBOARD DONUT =====================
function renderDashboardDonut(all) {
  const canvas = document.getElementById('status-donut');
  if (!canvas) return;
  const counts = { pending:0, posted:0, mailed:0, failed:0 };
  all.forEach(s => { if (counts[s.status] !== undefined) counts[s.status]++; });
  const total = all.length || 1;
  const segments = [
    { label:'Pending', pct: counts.pending/total, color:'#f59e0b', val: counts.pending },
    { label:'Posted',  pct: counts.posted/total,  color:'#6366f1', val: counts.posted  },
    { label:'Mailed',  pct: counts.mailed/total,  color:'#10b981', val: counts.mailed  },
    { label:'Failed',  pct: counts.failed/total,  color:'#ef4444', val: counts.failed  },
  ];
  const dpr = window.devicePixelRatio || 1;
  const SIZE = 200;
  canvas.width  = SIZE * dpr; canvas.height = SIZE * dpr;
  canvas.style.width = SIZE + 'px'; canvas.style.height = SIZE + 'px';
  const ctx = canvas.getContext('2d');
  ctx.scale(dpr, dpr);
  const cx = SIZE/2, cy = SIZE/2, R = 80, innerR = 50;
  if (all.length === 0) {
    ctx.beginPath(); ctx.arc(cx, cy, R, 0, Math.PI*2);
    ctx.fillStyle = 'rgba(255,255,255,0.05)'; ctx.fill();
  } else {
    let angle = -Math.PI / 2;
    segments.forEach(seg => {
      if (seg.val === 0) return;
      const sweep = seg.pct * Math.PI * 2;
      ctx.beginPath(); ctx.moveTo(cx, cy);
      ctx.arc(cx, cy, R, angle, angle + sweep);
      ctx.closePath(); ctx.fillStyle = seg.color; ctx.fill();
      angle += sweep;
    });
  }
  ctx.beginPath(); ctx.arc(cx, cy, innerR, 0, Math.PI * 2);
  ctx.fillStyle = getComputedStyle(document.body).getPropertyValue('--bg-primary') || '#060c1c';
  ctx.fill();
  ctx.textAlign = 'center'; ctx.textBaseline = 'middle';
  ctx.fillStyle = 'rgba(255,255,255,0.9)';
  ctx.font = `bold 22px 'Space Grotesk', Inter, sans-serif`;
  ctx.fillText(all.length, cx, cy - 8);
  ctx.font = `11px Inter, sans-serif`; ctx.fillStyle = 'rgba(255,255,255,0.4)';
  ctx.fillText('packages', cx, cy + 10);
  const legend = document.getElementById('status-legend');
  if (legend) {
    legend.innerHTML = segments.map(s => `
      <div class="legend-item" onclick="showPackagesByStatus('${s.label.toLowerCase()}')" style="cursor:pointer">
        <div class="legend-dot" style="background:${s.color}"></div>
        <span>${s.label}</span>
        <span class="legend-pct" style="color:${s.color}">${s.val}</span>
      </div>`).join('');
  }
}

// ===================== ACTIVITY LIST =====================
function renderActivityList(all) {
  const container = document.getElementById('activity-list');
  if (!container) return;
  if (!all.length) {
    container.innerHTML = '<div class="activity-empty">No shipments yet. Create your first shipment →</div>';
    return;
  }
  const sorted = [...all].sort((a,b) => new Date(b.created_at) - new Date(a.created_at)).slice(0,6);
  container.innerHTML = sorted.map(s => {
    const m   = statusMeta[s.status] || statusMeta.pending;
    const ico = itemMeta[s.item_type]?.icon || '📦';
    const t   = formatRelative(s.created_at);
    return `
      <div class="activity-item" onclick="openShipmentModal(${s.id})">
        <div class="activity-icon">${ico}</div>
        <div class="activity-info">
          <div class="activity-name">${s.item_desc}</div>
          <div class="activity-meta">${s.tracking} · ${t}</div>
        </div>
        <div class="activity-status" style="color:${m.color};background:${m.bg}">${m.icon} ${m.label.split(' /')[0]}</div>
      </div>`;
  }).join('');
}

// ===================== CATEGORY BARS =====================
function renderCategoryBars(all) {
  const container = document.getElementById('category-bars');
  if (!container) return;
  if (!all.length) {
    container.innerHTML = '<div class="activity-empty">No data yet.</div>';
    return;
  }
  const counts = {};
  all.forEach(s => { const key = s.item_type || 'other'; counts[key] = (counts[key] || 0) + 1; });
  const max = Math.max(...Object.values(counts));
  container.innerHTML = Object.entries(counts)
    .sort((a,b) => b[1]-a[1])
    .map(([type, cnt]) => {
      const m   = itemMeta[type] || itemMeta.other;
      const pct = (cnt / max) * 100;
      return `
        <div class="hbar-row">
          <div class="hbar-label">${m.icon} ${m.label.split(' /')[0]}</div>
          <div class="hbar-track"><div class="hbar-fill" style="width:${pct}%"></div></div>
          <div class="hbar-val">${cnt}</div>
        </div>`;
    }).join('');
}

// ===================== PACKAGES GRID =====================
let _currentFilter = 'all';
let _cachedPackages = [];

function setPackageFilter(filter, btn) {
  _currentFilter = filter;
  document.querySelectorAll('.filter-tab').forEach(b => b.classList.remove('active'));
  btn.classList.add('active');
  renderPackages();
}

function showPackagesByStatus(status) {
  showPage('packages');
  _currentFilter = status;
  document.querySelectorAll('.filter-tab').forEach(b => {
    b.classList.toggle('active', b.dataset.filter === status || (status === 'all' && b.dataset.filter === 'all'));
  });
  renderPackages();
}

async function renderPackages(preloaded) {
  const container = document.getElementById('packages-grid');
  if (!container) return;

  let all;
  if (preloaded) {
    _cachedPackages = preloaded;
    all = preloaded;
  } else {
    try {
      const result = await apiGet('/shipments');
      all = result.ok ? result.data : _cachedPackages;
      _cachedPackages = all;
    } catch (e) {
      all = _cachedPackages;
    }
  }

  const query = (document.getElementById('pkg-search')?.value || '').toLowerCase().trim();
  if (_currentFilter !== 'all') all = all.filter(s => s.status === _currentFilter);
  if (query) all = all.filter(s =>
    (s.item_desc||'').toLowerCase().includes(query) ||
    (s.tracking||'').toLowerCase().includes(query) ||
    (s.sender_name||'').toLowerCase().includes(query) ||
    (s.dest||'').toLowerCase().includes(query)
  );

  if (!all.length) {
    container.innerHTML = `<div class="packages-empty"><div style="font-size:48px;margin-bottom:12px">📬</div><div>No packages found.</div></div>`;
    return;
  }

  const sorted = [...all].sort((a,b) => new Date(b.created_at) - new Date(a.created_at));
  container.innerHTML = sorted.map(s => {
    const m   = statusMeta[s.status] || statusMeta.pending;
    const ico = itemMeta[s.item_type]?.icon || '📦';
    const t   = formatRelative(s.created_at);
    return `
      <div class="pkg-card" onclick="openShipmentModal(${s.id})">
        <div class="pkg-card-top">
          <div class="pkg-icon">${ico}</div>
          <div class="pkg-status-badge" style="color:${m.color};background:${m.bg}">${m.icon} ${s.status.charAt(0).toUpperCase()+s.status.slice(1)}</div>
        </div>
        <div class="pkg-name">${s.item_desc}</div>
        <div class="pkg-tracking">${s.tracking}</div>
        <div class="pkg-meta">
          <span>👤 ${s.sender_name}</span>
          <span>📍 ${s.dest}</span>
        </div>
        <div class="pkg-footer">
          <span class="pkg-courier">${s.courier.toUpperCase()} · ${s.pkg_size}</span>
          <span class="pkg-total">£${s.total?.toFixed(2)}</span>
        </div>
        <div class="pkg-date">${t}</div>
      </div>`;
  }).join('');
}

// ===================== SHIPMENT MODAL =====================
async function openShipmentModal(id) {
  const result = await apiGet(`/shipments/${id}`);
  if (!result.ok) return;
  const s = result.data;

  const m   = statusMeta[s.status] || statusMeta.pending;
  const ico = itemMeta[s.item_type]?.icon || '📦';

  document.getElementById('modal-tracking').textContent = s.tracking;
  document.getElementById('modal-item').textContent = s.item_desc;

  const body = document.getElementById('modal-body');
  body.innerHTML = `
    <div class="modal-detail-grid">
      <div class="modal-icon-wrap">${ico}</div>
      <div class="modal-details">
        <div class="modal-status-big" style="color:${m.color}">${m.icon} ${m.label}</div>
        <div class="detail-row"><span>Sender</span><strong>${s.sender_name}</strong></div>
        <div class="detail-row"><span>Email</span><strong>${s.sender_email || '—'}</strong></div>
        <div class="detail-row"><span>Destination</span><strong>${s.dest}</strong></div>
        <div class="detail-row"><span>Address</span><strong>${s.recv_address || '—'}, ${s.recv_city} ${s.recv_postcode}</strong></div>
        <div class="detail-row"><span>Courier</span><strong>${s.courier?.toUpperCase()} · ${s.pkg_size}</strong></div>
        <div class="detail-row"><span>Total Paid</span><strong>£${s.total?.toFixed(2)}</strong></div>
        <div class="detail-row"><span>Created</span><strong>${formatDate(s.created_at)}</strong></div>
      </div>
    </div>
    <div class="modal-timeline">
      <div class="modal-timeline-title">Status History</div>
      ${(s.timeline||[]).map((e,i) => `
        <div class="mtl-item ${i===s.timeline.length-1?'active':'done'}">
          <div class="mtl-dot"></div>
          <div class="mtl-content">
            <div class="mtl-status">${statusMeta[e.status]?.label || e.status}</div>
            ${e.note ? `<div class="mtl-note">${e.note}</div>` : ''}
            <div class="mtl-time">${formatDate(e.ts)}</div>
          </div>
        </div>`).join('')}
    </div>`;

  const statuses = ['pending','posted','mailed','failed'];
  const btnArea  = document.getElementById('modal-status-btns');
  btnArea.innerHTML = statuses.map(st => {
    const active = st === s.status ? 'active-status-btn' : '';
    const sm = statusMeta[st];
    return `<button class="status-change-btn ${active}" style="${active ? `background:${sm.bg};border-color:${sm.color};color:${sm.color}` : ''}" onclick="changeStatus(${s.id},'${st}')">${sm.icon} ${st.charAt(0).toUpperCase()+st.slice(1)}</button>`;
  }).join('');
  btnArea.innerHTML += `<button class="status-change-btn delete-btn" onclick="deleteShipment(${s.id})" style="border-color:var(--danger);color:var(--danger)">🗑 Delete</button>`;

  document.getElementById('modal-overlay').classList.add('open');
  document.getElementById('shipment-modal').classList.add('open');
}

function closeModal() {
  document.getElementById('modal-overlay').classList.remove('open');
  document.getElementById('shipment-modal').classList.remove('open');
}

async function changeStatus(id, newStatus) {
  const result = await apiPut(`/shipments/${id}/status`, { status: newStatus });
  if (result.ok) {
    closeModal();
    await refreshAll();
    showToast(`Status updated → ${statusMeta[newStatus].label}`, 'success');
  } else {
    showToast('Failed to update status', 'error');
  }
}

async function deleteShipment(id) {
  if (!confirm('Delete this shipment from the database?')) return;
  const result = await apiDelete(`/shipments/${id}`);
  if (result.ok) {
    closeModal();
    await refreshAll();
    showToast('Shipment deleted.', 'info');
  } else {
    showToast('Failed to delete', 'error');
  }
}

// ===================== TRACKING =====================
async function trackPackage() {
  const input = document.getElementById('track-input').value.trim();
  if (!input) { showToast('Please enter a tracking number.', 'error'); return; }

  try {
    const result = await apiGet(`/shipments/track/${encodeURIComponent(input.toUpperCase())}`);
    const resultEl = document.getElementById('tracking-result');
    if (!result.ok) {
      showToast(`No shipment found for ${input}`, 'error');
      resultEl.style.display = 'none';
      return;
    }
    const s = result.data;
    resultEl.style.display = 'block';
    document.getElementById('t-id').textContent   = s.tracking;
    document.getElementById('t-item').textContent = s.item_desc;
    const m  = statusMeta[s.status] || statusMeta.pending;
    const sb = document.getElementById('t-status');
    sb.textContent = m.label;
    sb.style.background = m.bg;
    sb.style.color      = m.color;
    sb.className = 'badge';
    const tl = document.getElementById('track-timeline');
    tl.innerHTML = (s.timeline || []).map((e,i) => {
      const isDone   = i < s.timeline.length - 1;
      const isActive = i === s.timeline.length - 1;
      return `
        <div class="timeline-item ${isDone?'done':''} ${isActive?'active':''}">
          <div class="tl-dot"></div>
          <div class="tl-content">
            <div class="tl-title">${statusMeta[e.status]?.label || e.status}</div>
            <div class="tl-time">${formatDate(e.ts)}</div>
          </div>
        </div>`;
    }).join('');
    resultEl.scrollIntoView({ behavior:'smooth', block:'nearest' });
    showToast(`Tracking loaded for ${input}`, 'info');
  } catch (e) {
    showToast('Cannot reach server', 'error');
  }
}

// Quick track from user mode landing
function quickTrack() {
  const input = document.getElementById('quick-track-input')?.value?.trim();
  if (!input) { showToast('Enter a tracking number', 'error'); return; }
  // Enter admin mode first if needed, or show public track result
  if (currentMode === MODE.USER) {
    // Show inline result in user mode
    document.getElementById('user-track-result').style.display = 'block';
    const trackInput = document.getElementById('track-input');
    if (trackInput) trackInput.value = input;
    trackPackagePublic(input);
  }
}

async function trackPackagePublic(code) {
  const container = document.getElementById('public-track-result');
  if (!container) return;
  container.innerHTML = '<div class="tracking-loading">Searching…</div>';
  container.style.display = 'block';
  try {
    const result = await apiGet(`/shipments/track/${encodeURIComponent(code.toUpperCase())}`);
    if (!result.ok) {
      container.innerHTML = `<div class="public-track-error">❌ No shipment found for <strong>${code}</strong></div>`;
      return;
    }
    const s = result.data;
    const m = statusMeta[s.status] || statusMeta.pending;
    const ico = itemMeta[s.item_type]?.icon || '📦';
    container.innerHTML = `
      <div class="public-track-card">
        <div class="public-track-header">
          <div class="public-track-icon">${ico}</div>
          <div>
            <div class="public-track-id">${s.tracking}</div>
            <div class="public-track-desc">${s.item_desc}</div>
          </div>
          <div class="public-track-status" style="color:${m.color};background:${m.bg}">${m.icon} ${m.label}</div>
        </div>
        <div class="public-track-timeline">
          ${(s.timeline||[]).map((e,i) => `
            <div class="pub-tl-item ${i===s.timeline.length-1?'active':'done'}">
              <div class="pub-tl-dot"></div>
              <div class="pub-tl-text">
                <div style="font-weight:600">${statusMeta[e.status]?.label||e.status}</div>
                <div style="opacity:.6;font-size:.8rem">${formatDate(e.ts)}</div>
              </div>
            </div>`).join('')}
        </div>
        <div style="font-size:.8rem;opacity:.5;margin-top:1rem">Destination: ${s.dest}</div>
      </div>`;
  } catch (e) {
    container.innerHTML = '<div class="public-track-error">⚠️ Server offline</div>';
  }
}

// ===================== EXPORT CSV =====================
async function exportCSV() {
  let result = await apiGet('/shipments');
  let all = result.ok ? result.data : [];
  if (_currentFilter !== 'all') all = all.filter(s => s.status === _currentFilter);
  const headers = ['Tracking','Item','Sender','Email','Destination','Courier','Package','Status','Total','Currency','Created'];
  const rows    = all.map(s => [
    s.tracking, s.item_desc, s.sender_name, s.sender_email||'',
    s.dest, s.courier?.toUpperCase(), s.pkg_size, s.status,
    formatMoney(s.total || 0), activeCurrency, formatDate(s.created_at),
  ]);
  const csv  = [headers, ...rows].map(r => r.map(v => `"${String(v).replace(/"/g,'""')}"`).join(',')).join('\n');
  const blob = new Blob([csv], { type: 'text/csv' });
  const url  = URL.createObjectURL(blob);
  const a    = Object.assign(document.createElement('a'), { href:url, download:`sendit_${_currentFilter}_${new Date().toISOString().slice(0,10)}.csv` });
  a.click();
  URL.revokeObjectURL(url);
  showToast('CSV exported!', 'success');
}

// ===================== STEP NAVIGATION =====================
function nextStep(step) {
  if (step === 2 && state.currentStep === 1) {
    if (!document.getElementById('item-type').value || !document.getElementById('item-desc').value.trim()) {
      showToast('Please select an item type and provide a description.', 'error');
      return;
    }
  }
  if (step === 3 && state.currentStep === 2) {
    const required = ['sender-name','recv-name','recv-address1','recv-city','recv-postcode'];
    if (required.some(id => !document.getElementById(id).value.trim())) {
      showToast('Please fill all required address fields.', 'error');
      return;
    }
  }
  document.getElementById(`form-step-${state.currentStep}`).classList.remove('active');
  document.getElementById(`step-${state.currentStep}`).classList.remove('active');
  document.getElementById(`step-${state.currentStep}`).classList.add('done');
  state.currentStep = step;
  document.getElementById(`form-step-${step}`).classList.add('active');
  for (let i=1; i<step; i++) { document.getElementById(`step-${i}`).classList.add('done'); document.getElementById(`step-${i}`).classList.remove('active'); }
  for (let i=step; i<=5; i++) { document.getElementById(`step-${i}`).classList.remove('done'); }
  document.getElementById(`step-${step}`).classList.add('active');
  if (step === 4) updateTotals();
}

// ===================== ITEM ICON =====================
function updateItemIcon() {
  const sel  = document.getElementById('item-type').value;
  const meta = itemMeta[sel];
  document.getElementById('item-icon-display').textContent = meta ? meta.icon : '📦';
  document.getElementById('item-preview-name').textContent = meta ? meta.label : 'Select an item type';
  const be = document.getElementById('badge-eligible');
  const bh = document.getElementById('badge-hazmat');
  if (meta) {
    be.textContent = 'Yes ✓'; be.className = 'badge badge-ok';
    bh.textContent = meta.hazmat; bh.className = meta.hazmat === 'None' ? 'badge badge-ok' : 'badge badge-warn';
  } else {
    be.textContent = 'Pending'; be.className = 'badge badge-pending';
    bh.textContent = '—'; bh.className = 'badge badge-pending';
  }
}

// ===================== PACKAGE SELECTION =====================
function selectPackage(size, el) {
  document.querySelectorAll('.package-option').forEach(o => o.classList.remove('selected'));
  el.classList.add('selected');
  state.selectedPackage = size;
  updateCourierRates();
  updateTotals();
}
function updateCourierRates() {
  const sz = state.selectedPackage;
  setText('dhl-rate', formatMoney(pricing.couriers.dhl[sz]));
  setText('dpd-rate', formatMoney(pricing.couriers.dpd[sz]));
  setText('rm-rate',  formatMoney(pricing.couriers.rm[sz]));
}

// ===================== TOTALS =====================
function toggleInsurance() {
  state.insuranceAdded = document.getElementById('insurance-toggle').checked;
  updateTotals();
}
function updateTotals() {
  const postage = pricing.couriers[state.selectedCourier][state.selectedPackage];
  const pkgFee  = pricing.packages[state.selectedPackage];
  const ins     = state.insuranceAdded ? pricing.insurance : 0;
  const total   = postage + pkgFee + pricing.serviceFee + ins;
  const cName   = { dhl:'DHL Express', dpd:'DPD Tracked', rm:'Royal Mail 1st' }[state.selectedCourier];
  const pName   = { envelope:'Envelope', small:'Small Box', medium:'Medium Box' }[state.selectedPackage];
  setText('lbl-postage', `Postage (${cName})`);
  setText('lbl-package', `Package (${pName})`);
  setText('sum-postage', formatMoney(postage));
  setText('sum-package', formatMoney(pkgFee));
  setText('sum-fee',     formatMoney(pricing.serviceFee));
  setText('sum-total',   formatMoney(total));
}

// ===================== PAYMENT =====================
async function processPayment() {
  const cardNum = document.getElementById('card-number').value.trim();
  if (!cardNum && !state.selectedPayMethod) {
    showToast('Please enter payment details or select a payment method.', 'error');
    return;
  }
  const btn = document.querySelector('.btn-pay');
  btn.textContent = 'Saving to database…';
  btn.disabled = true;
  await new Promise(r => setTimeout(r, 1200));

  const postage = pricing.couriers[state.selectedCourier][state.selectedPackage];
  const pkgFee  = pricing.packages[state.selectedPackage];
  const total   = postage + pkgFee + pricing.serviceFee + (state.insuranceAdded ? pricing.insurance : 0);

  const record = {
    item_type:    document.getElementById('item-type').value,
    item_desc:    document.getElementById('item-desc').value.trim(),
    item_value:   parseFloat(document.getElementById('item-value').value) || 0,
    reason:       document.querySelector('input[name="reason"]:checked')?.value || 'confiscated',
    sender_name:  document.getElementById('sender-name').value.trim(),
    sender_email: document.getElementById('sender-email').value.trim(),
    sender_phone: document.getElementById('sender-phone').value.trim(),
    recv_name:    document.getElementById('recv-name').value.trim(),
    recv_address: document.getElementById('recv-address1').value.trim() + (document.getElementById('recv-address2').value.trim() ? ', ' + document.getElementById('recv-address2').value.trim() : ''),
    recv_city:    document.getElementById('recv-city').value.trim(),
    recv_postcode:document.getElementById('recv-postcode').value.trim(),
    recv_country: document.getElementById('recv-country').value,
    dest:         `${document.getElementById('recv-city').value.trim()}, ${document.getElementById('recv-country').value}`,
    courier:      state.selectedCourier,
    pkg_size:     state.selectedPackage,
    total,
    insurance:    state.insuranceAdded,
  };

  try {
    const result = await apiPost('/shipments', record);
    btn.textContent = 'Pay & Save to Database';
    btn.disabled = false;
    if (result.ok) {
      setText('confirm-tracking', `Tracking: ${result.data.tracking}`);
      nextStep(5);
      await refreshAll();
      showToast(`Shipment ${result.data.tracking} saved to database!`, 'success');
    } else {
      showToast('Failed to save: ' + result.error, 'error');
    }
  } catch (e) {
    btn.textContent = 'Pay & Save to Database';
    btn.disabled = false;
    showToast('Server error: ' + e.message, 'error');
  }
}

// ===================== QR SESSION (One-Time Code) =====================
let qrSessionCode       = null;
let qrSessionPollTimer  = null;
let qrSessionCountTimer = null;

async function openQRSession() {
  const qr = document.getElementById('qr-box');
  qr.innerHTML = `<div class="qr-loading">Generating code…</div>`;

  try {
    const resp = await fetch(`${API.replace('/api','')}/api/session/create`, { method:'POST' });
    const data = await resp.json();
    if (!data.ok) throw new Error(data.error);

    qrSessionCode = data.code;
    const mobileUrl = data.mobile_url;
    const expiresAt = new Date(data.expires_at);
    const qrImgUrl  = `https://api.qrserver.com/v1/create-qr-code/?size=160x160&data=${encodeURIComponent(mobileUrl)}&bgcolor=ffffff&color=1a1a2e&qzone=1`;

    qr.innerHTML = `
      <div class="qr-session-wrap">
        <div class="qr-session-left">
          <img src="${qrImgUrl}" width="160" height="160" alt="QR Code" class="qr-img" />
          <div class="qr-or">or type this URL:</div>
          <div class="qr-url">${mobileUrl.replace(/https?:\/\//,'')}</div>
        </div>
        <div class="qr-session-right">
          <div class="qr-code-label">Your one-time code</div>
          <div class="qr-code-display" id="qr-code-display">${data.code}</div>
          <div class="qr-code-hint">Scan the QR <em>or</em> go to the URL on your phone, then enter this code.</div>
          <div class="qr-timer" id="qr-timer">Expires in 10:00</div>
          <button class="qr-cancel-btn" onclick="cancelQRSession()">Cancel</button>
        </div>
      </div>`;

    // Countdown timer display
    const totalMs = expiresAt - Date.now();
    qrSessionCountTimer = setInterval(() => {
      const left = Math.max(0, Math.floor((expiresAt - Date.now()) / 1000));
      const m = Math.floor(left / 60), s = left % 60;
      const el = document.getElementById('qr-timer');
      if (el) el.textContent = `Expires in ${m}:${String(s).padStart(2,'0')}`;
      if (left === 0) cancelQRSession();
    }, 1000);

    // Poll for data
    startQRSessionPoll(data.code);
    showToast('Show the QR or code to the passenger', 'info');
  } catch(e) {
    qr.innerHTML = `<span style="color:var(--danger)">Failed to generate code — is the server running?</span>`;
  }
}

function startQRSessionPoll(code) {
  qrSessionPollTimer = setInterval(async () => {
    try {
      const resp = await fetch(`${API.replace('/api','')}/api/session/poll/${code}`);
      const data = await resp.json();
      if (data.ok && data.status === 'received') {
        clearInterval(qrSessionPollTimer);
        clearInterval(qrSessionCountTimer);
        autoFillAllFromMobileData(data.data);
        nextStep(4); // Skip straight to step 4 since everything is filled
      } else if (data.status === 'expired') {
        clearInterval(qrSessionPollTimer);
        clearInterval(qrSessionCountTimer);
        const qr = document.getElementById('qr-box');
        if (qr) qr.innerHTML = `<span style="color:var(--warning)">Code expired. Click to generate new one.</span>`;
        qr.onclick = openQRSession;
      }
    } catch(e) { /* ignore network blips */ }
  }, 2000);
}

function cancelQRSession() {
  clearInterval(qrSessionPollTimer);
  clearInterval(qrSessionCountTimer);
  qrSessionCode = null;
  const qr = document.getElementById('qr-box');
  if (qr) {
    qr.innerHTML = `<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/><rect x="5" y="5" width="3" height="3" fill="currentColor" stroke="none"/><rect x="16" y="5" width="3" height="3" fill="currentColor" stroke="none"/><rect x="5" y="16" width="3" height="3" fill="currentColor" stroke="none"/></svg><span>Scan QR from phone to auto-fill address</span>`;
    qr.onclick = openQRSession;
    qr.style.color = ''; qr.style.borderColor = '';
  }
}

function autoFillAllFromMobileData(d) {
  // Step 1: Item Details
  if (d.item_type) document.getElementById('item-type').value = d.item_type;
  if (d.item_desc) document.getElementById('item-desc').value = d.item_desc;
  if (d.item_value) document.getElementById('item-value').value = d.item_value;
  if (d.reason) {
    const radio = document.querySelector(`input[name="reason"][value="${d.reason}"]`);
    if (radio) radio.checked = true;
  }
  updateItemIcon();

  // Step 2: Address Details
  if (d.sender_name)  document.getElementById('sender-name').value  = d.sender_name;
  if (d.sender_email) document.getElementById('sender-email').value = d.sender_email;
  if (d.sender_phone) document.getElementById('sender-phone').value = d.sender_phone;
  if (d.recv_name)    document.getElementById('recv-name').value    = d.recv_name;
  if (d.recv_address1) document.getElementById('recv-address1').value = d.recv_address1;
  else if (d.recv_address) document.getElementById('recv-address1').value = d.recv_address;
  if (d.recv_address2) document.getElementById('recv-address2').value = d.recv_address2;
  if (d.recv_city)    document.getElementById('recv-city').value    = d.recv_city;
  if (d.recv_postcode)document.getElementById('recv-postcode').value = d.recv_postcode;
  if (d.recv_country) {
    const sel = document.getElementById('recv-country');
    if (sel) sel.value = d.recv_country;
  }

  // Step 3: Package size & Courier
  if (d.pkg_size) {
    state.selectedPackage = d.pkg_size;
    document.querySelectorAll('.package-option').forEach(o => o.classList.remove('selected'));
    const pkgEl = document.getElementById(`pkg-${d.pkg_size}`);
    if (pkgEl) pkgEl.classList.add('selected');
  }
  if (d.courier) {
    state.selectedCourier = d.courier;
    const radio = document.querySelector(`input[name="courier"][value="${d.courier}"]`);
    if (radio) radio.checked = true;
    document.querySelectorAll('.courier-option').forEach(o => o.classList.remove('active'));
    const map = { dhl:'courier-dhl', dpd:'courier-dpd', rm:'courier-rm' };
    const courierEl = document.getElementById(map[d.courier]);
    if (courierEl) courierEl.classList.add('active');
  }

  // Insurance
  if (d.insurance !== undefined) {
    state.insuranceAdded = !!d.insurance;
    const toggle = document.getElementById('insurance-toggle');
    if (toggle) toggle.checked = state.insuranceAdded;
  }

  // Step 4: Payment Details
  if (d.card_number) document.getElementById('card-number').value = d.card_number;
  if (d.card_expiry) document.getElementById('card-expiry').value = d.card_expiry;
  if (d.card_cvv)    document.getElementById('card-cvv').value    = d.card_cvv;
  if (d.card_name)   document.getElementById('card-name').value   = d.card_name;
  if (d.pay_method) {
    state.selectedPayMethod = d.pay_method;
    document.querySelectorAll('.pay-btn').forEach(b => b.classList.remove('selected'));
    const payBtn = document.getElementById(`pay-${d.pay_method}`);
    if (payBtn) payBtn.classList.add('selected');
  } else {
    state.selectedPayMethod = null;
    document.querySelectorAll('.pay-btn').forEach(b => b.classList.remove('selected'));
  }

  updateCourierRates();
  updateTotals();

  const qr = document.getElementById('qr-box');
  if (qr) {
    qr.innerHTML = `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg><span>✓ Form auto-filled from passenger's phone!</span>`;
    qr.style.borderColor = 'var(--success)'; qr.style.color = 'var(--success)';
    qr.onclick = null;
  }
  showToast('🎉 Passenger data received! Form auto-filled.', 'success');
}

function autoFillFromMobileData(d) {
  autoFillAllFromMobileData(d);
}

function selectPayMethod(method) {
  state.selectedPayMethod = method;
  document.querySelectorAll('.pay-btn').forEach(b => b.classList.remove('selected'));
  document.getElementById(`pay-${method}`).classList.add('selected');
  document.getElementById('card-number').value = '';
  showToast(`${method.charAt(0).toUpperCase()+method.slice(1)} Pay selected`, 'info');
}

function formatCard(el) {
  let v = el.value.replace(/\D/g,'').substring(0,16);
  el.value = v.match(/.{1,4}/g)?.join(' ') || v;
}
function formatExpiry(el) {
  let v = el.value.replace(/\D/g,'').substring(0,4);
  if (v.length >= 2) v = v.substring(0,2) + '/' + v.substring(2);
  el.value = v;
}

function printLabel() {
  showToast('Sending to label printer…', 'success');
}

function resetForm() {
  clearInterval(qrSessionPollTimer);
  clearInterval(qrSessionCountTimer);
  qrSessionCode = null;
  document.querySelectorAll('#page-new-shipment input, #page-new-shipment select').forEach(el => {
    if (el.type === 'checkbox') el.checked = false;
    else if (el.type === 'radio') el.checked = el.value === 'confiscated';
    else el.value = '';
  });
  document.getElementById('item-icon-display').textContent = '📦';
  document.getElementById('item-preview-name').textContent = 'Select an item type';
  const qr = document.getElementById('qr-box');
  qr.style.color = ''; qr.style.borderColor = '';
  qr.innerHTML = '<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/><rect x="5" y="5" width="3" height="3" fill="currentColor" stroke="none"/><rect x="16" y="5" width="3" height="3" fill="currentColor" stroke="none"/><rect x="5" y="16" width="3" height="3" fill="currentColor" stroke="none"/></svg><span>Scan QR or enter code to auto-fill from phone</span>';
  qr.onclick = openQRSession;
  document.querySelectorAll('.package-option').forEach(o => o.classList.remove('selected'));
  document.getElementById('pkg-small').classList.add('selected');
  state.selectedPackage = 'small'; state.insuranceAdded = false; state.selectedPayMethod = null;
  state.currentStep = 1;
  document.querySelectorAll('.form-step').forEach(s => s.classList.remove('active'));
  document.getElementById('form-step-1').classList.add('active');
  document.querySelectorAll('.step').forEach(s => { s.classList.remove('active','done'); });
  document.getElementById('step-1').classList.add('active');
  updateTotals();
}

// ===================== NAVIGATION =====================
function showPage(pageId) {
  document.querySelectorAll('.page').forEach(p => p.classList.remove('active'));
  document.querySelectorAll('.nav-item').forEach(n => n.classList.remove('active'));
  const pageMap = {
    'dashboard':    { dom:'page-dashboard',    nav:'nav-dashboard', title:'Dashboard',      crumb:'SENDIT → Overview' },
    'new-shipment': { dom:'page-new-shipment', nav:'nav-new',       title:'New Shipment',   crumb:'SENDIT → New Shipment' },
    'packages':     { dom:'page-packages',     nav:'nav-packages',  title:'All Packages',   crumb:'SENDIT → All Packages' },
    'track':        { dom:'page-track',        nav:'nav-track',     title:'Track Package',  crumb:'SENDIT → Track' },
    'rates':        { dom:'page-rates',        nav:'nav-rates',     title:'Rates & Sizes',  crumb:'SENDIT → Rates' },
  };
  const p = pageMap[pageId]; if (!p) return;
  document.getElementById(p.dom).classList.add('active');
  document.getElementById(p.nav)?.classList.add('active');
  setText('page-title', p.title);
  setText('breadcrumb',  p.crumb);
  state.currentPage = pageId;
  if (pageId === 'packages')  renderPackages();
  if (pageId === 'dashboard') refreshAll();
  if (currentMode === MODE.ADMIN) resetAdminTimeout();
}

// ===================== SETTINGS =====================
function openSettings() { document.getElementById('settings-overlay').classList.add('open'); document.getElementById('settings-panel').classList.add('open'); }
function closeSettings() { document.getElementById('settings-overlay').classList.remove('open'); document.getElementById('settings-panel').classList.remove('open'); }

function setTheme(t, btn) {
  const b = document.getElementById('app-body');
  ['dark-blue','dark-green','black','white','deep-purple','slate'].forEach(x => b.classList.remove(`theme-${x}`));
  b.classList.add(`theme-${t}`); state.theme = t;
  document.querySelectorAll('.theme-btn').forEach(x => x.classList.remove('active'));
  btn.classList.add('active');
  savePreferences();
  apiGet('/shipments').then(r => r.ok && renderDashboardDonut(r.data));
}
function setTextColor(c, btn) {
  const b = document.getElementById('app-body');
  ['white','blue-soft','green-soft','amber','rose','slate-dark'].forEach(x => b.classList.remove(`text-${x}`));
  b.classList.add(`text-${c}`); state.textColor = c;
  document.querySelectorAll('.color-btn[data-color]').forEach(x => x.classList.remove('active'));
  btn.classList.add('active'); savePreferences();
}
function setAccent(a, btn) {
  const b = document.getElementById('app-body');
  ['cyan','indigo','emerald','orange','pink','gold'].forEach(x => b.classList.remove(`accent-${x}`));
  b.classList.add(`accent-${a}`); state.accent = a;
  document.querySelectorAll('.color-btn[data-accent]').forEach(x => x.classList.remove('active-accent'));
  btn.classList.add('active-accent'); savePreferences();
}
function setFontSize(s) {
  const b = document.getElementById('app-body');
  ['small','medium','large'].forEach(x => b.classList.remove(`font-${x}`));
  b.classList.add(`font-${s}`); state.fontSize = s;
  document.querySelectorAll('.font-btn').forEach(x => x.classList.remove('active'));
  document.querySelector(`.font-btn[onclick*="${s}"]`)?.classList.add('active');
  savePreferences();
}
function savePreferences() {
  try { localStorage.setItem('sendit_prefs', JSON.stringify({ theme:state.theme, textColor:state.textColor, accent:state.accent, fontSize:state.fontSize, currency: activeCurrency })); } catch(e) {}
}
function loadPreferences() {
  try {
    const p = JSON.parse(localStorage.getItem('sendit_prefs') || '{}');
    if (p.theme)     setTheme(p.theme, document.querySelector(`[data-theme="${p.theme}"]`) || document.querySelector('.theme-btn'));
    if (p.textColor) setTextColor(p.textColor, document.querySelector(`[data-color="${p.textColor}"]`) || document.querySelector('.color-btn[data-color]'));
    if (p.accent)    setAccent(p.accent, document.querySelector(`[data-accent="${p.accent}"]`) || document.querySelector('.color-btn[data-accent]'));
    if (p.fontSize)  setFontSize(p.fontSize);
    if (p.currency && CURRENCIES[p.currency]) activeCurrency = p.currency;
  } catch(e) {}
}

// ===================== PIN SETTINGS =====================
async function changePIN() {
  const current = document.getElementById('current-pin-input')?.value?.trim();
  const newPin  = document.getElementById('new-pin-input')?.value?.trim();
  if (!current || !newPin) { showToast('Fill both PIN fields', 'error'); return; }
  const result = await apiPut('/settings/pin', { current_pin: current, new_pin: newPin });
  if (result.ok) {
    showToast('PIN changed successfully!', 'success');
    document.getElementById('current-pin-input').value = '';
    document.getElementById('new-pin-input').value = '';
  } else {
    showToast(result.error || 'Failed to change PIN', 'error');
  }
}

// ===================== CLOCK =====================
function startClock() {
  const up = () => {
    const n  = new Date();
    const el = document.getElementById('live-time');
    if (el) el.textContent = n.toLocaleDateString('en-GB',{day:'2-digit',month:'short',year:'numeric'}) + ' · ' + n.toLocaleTimeString('en-GB');
  };
  up(); setInterval(up, 1000);
}
function updateQueueCount() {
  const el = document.getElementById('queue-count');
  if (el) el.textContent = Math.max(0, Math.min(12, parseInt(el.textContent) + (Math.floor(Math.random()*3)-1)));
}

// ===================== UTILITIES =====================
function setText(id, val) { const e = document.getElementById(id); if (e) e.textContent = val; }

function formatDate(iso) {
  try { return new Date(iso).toLocaleString('en-GB', { day:'2-digit', month:'short', year:'numeric', hour:'2-digit', minute:'2-digit' }); }
  catch(e) { return iso; }
}
function formatRelative(iso) {
  try {
    const diff = (Date.now() - new Date(iso)) / 1000;
    if (diff < 60)   return 'just now';
    if (diff < 3600) return `${Math.floor(diff/60)}m ago`;
    if (diff < 86400)return `${Math.floor(diff/3600)}h ago`;
    return `${Math.floor(diff/86400)}d ago`;
  } catch(e) { return iso; }
}

// ===================== TOAST =====================
function showToast(msg, type='info') {
  const icons = {
    success:`<svg class="toast-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>`,
    error:  `<svg class="toast-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>`,
    info:   `<svg class="toast-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>`,
  };
  const c = document.getElementById('toast-container');
  const t = document.createElement('div');
  t.className = `toast toast-${type}`;
  t.innerHTML = `${icons[type]||icons.info}<span>${msg}</span>`;
  c.appendChild(t);
  setTimeout(() => { t.style.animation='toastOut 0.3s ease forwards'; setTimeout(()=>t.remove(),300); }, 3500);
}

// ===================== ENTER CODE (Pre-fill from Phone) =====================
let ecBuffer = '';

function openEnterCodeOverlay() {
  ecBuffer = '';
  updateEcDisplay();
  document.getElementById('ec-code-error').textContent = '';
  document.getElementById('enter-code-overlay').classList.add('open');
}

function closeEnterCodeOverlay(e) {
  if (e && e.target === document.getElementById('enter-code-overlay')) {
    closeEnterCodeOverlayDirect();
  }
}
function closeEnterCodeOverlayDirect() {
  document.getElementById('enter-code-overlay').classList.remove('open');
  ecBuffer = '';
  updateEcDisplay();
  document.getElementById('ec-code-error').textContent = '';
}

function ecKey(char) {
  if (ecBuffer.length >= 6) return;
  ecBuffer += char;
  updateEcDisplay();
  document.getElementById('ec-code-error').textContent = '';
  if (ecBuffer.length === 6) {
    // Auto-submit after slight delay for UX feedback
    setTimeout(submitEnterCode, 250);
  }
}

function ecBackspace() {
  ecBuffer = ecBuffer.slice(0, -1);
  updateEcDisplay();
  document.getElementById('ec-code-error').textContent = '';
}

function updateEcDisplay() {
  const el = document.getElementById('ec-code-display');
  if (!el) return;
  const chars = ecBuffer.split('');
  const display = Array.from({length: 6}, (_, i) => chars[i] || '_').join(' ');
  el.textContent = display;
}

let _ecPendingData = null;

async function submitEnterCode() {
  if (ecBuffer.length < 6) {
    document.getElementById('ec-code-error').textContent = 'Please enter the full 6-character code.';
    return;
  }

  const btn = document.getElementById('ec-submit-btn');
  if (btn) { btn.textContent = '⏳ Loading…'; btn.disabled = true; }

  try {
    const resp = await fetch(`${API}/prefill/redeem/${encodeURIComponent(ecBuffer)}`, { method: 'POST' });
    const data = await resp.json();

    if (btn) { btn.innerHTML = '<svg width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg> Confirm & Load My Data'; btn.disabled = false; }

    if (data.ok) {
      _ecPendingData = data.data;
      closeEnterCodeOverlayDirect();
      showEcConfirmOverlay(data.data);
    } else if (resp.status === 410) {
      document.getElementById('ec-code-error').textContent = '⏱ This code has expired. Codes are valid for 3 hours.';
    } else if (resp.status === 409) {
      document.getElementById('ec-code-error').textContent = '✓ This code has already been used at a kiosk.';
    } else if (resp.status === 404) {
      document.getElementById('ec-code-error').textContent = '✗ Code not found. Check the code and try again.';
    } else {
      document.getElementById('ec-code-error').textContent = data.error || 'Something went wrong.';
    }
  } catch(e) {
    if (btn) { btn.innerHTML = '<svg width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg> Confirm & Load My Data'; btn.disabled = false; }
    document.getElementById('ec-code-error').textContent = 'Cannot reach server — check connection.';
  }
}

function showEcConfirmOverlay(d) {
  // Build summary
  const rows = [
    ['Item Type', d.item_type ? (d.item_type.charAt(0).toUpperCase() + d.item_type.slice(1)) : '—'],
    ['Item', d.item_desc || '—'],
    ['Sender', d.sender_name || '—'],
    ['Recipient', d.recv_name || '—'],
    ['Destination', d.recv_city ? `${d.recv_city}, ${d.recv_country || ''}` : '—'],
    ['Package Size', d.pkg_size ? d.pkg_size.charAt(0).toUpperCase() + d.pkg_size.slice(1) : '—'],
    ['Courier', d.courier ? d.courier.toUpperCase() : '—'],
    ['Total', d.total ? `£${parseFloat(d.total).toFixed(2)}` : '—'],
  ];

  const summary = document.getElementById('ec-confirm-summary');
  if (summary) {
    summary.innerHTML = rows.map(([label, val]) => `
      <div class="ec-summary-row">
        <span>${label}</span>
        <span>${val}</span>
      </div>`).join('');
  }

  document.getElementById('ec-confirm-overlay').classList.add('open');
}

function closeEcConfirmOverlay(e) {
  if (!e || e.target === document.getElementById('ec-confirm-overlay')) {
    document.getElementById('ec-confirm-overlay').classList.remove('open');
    _ecPendingData = null;
  }
}

function proceedFromEcConfirm() {
  if (!_ecPendingData) return;
  const d = _ecPendingData;
  _ecPendingData = null;

  document.getElementById('ec-confirm-overlay').classList.remove('open');

  // Enter kiosk self-service flow with pre-filled data
  resetForm();
  document.body.classList.add('user-kiosk-flow');
  currentMode = MODE.USER;
  document.getElementById('user-mode-screen').classList.remove('active');
  document.getElementById('admin-panel').classList.add('active');
  showPage('new-shipment');

  // Auto-fill all form fields
  autoFillAllFromMobileData(d);

  // Jump straight to payment step (step 4) — all data is filled
  nextStep(4);

  showToast('🎉 Your pre-filled data loaded! Review payment and confirm.', 'success');
}


// ===================== GENERATE MY CODE OVERLAY =====================
function openGenerateCodeOverlay() {
  // Build the mobile prefill URL dynamically (works on any host/port)
  const base = `${window.location.protocol}//${window.location.host}`;
  const mobileUrl = `${base}/mobile?mode=phone`;

  // Generate QR via free qrserver API
  const qrUrl = `https://api.qrserver.com/v1/create-qr-code/?size=220x220&data=${encodeURIComponent(mobileUrl)}&bgcolor=ffffff&color=0f4c75&qzone=2`;
  const img = document.getElementById('gencode-qr-img');
  if (img) img.src = qrUrl;

  const urlEl = document.getElementById('gencode-url');
  if (urlEl) urlEl.textContent = mobileUrl.replace(/https?:\/\//, '');

  document.getElementById('gencode-overlay').classList.add('open');
}

function closeGenerateCodeOverlay(e) {
  if (e && e.target === document.getElementById('gencode-overlay')) {
    closeGenerateCodeOverlayDirect();
  }
}

function closeGenerateCodeOverlayDirect() {
  document.getElementById('gencode-overlay').classList.remove('open');
}
