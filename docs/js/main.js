/* ============================================================
   gos1mple — Theme Switcher, Modal, and Mobile Nav
   ============================================================ */

(function () {
  'use strict';

  const THEME_KEY = 'gos1mple-theme';

  // ── Multi-tier Theme Resolution ─────────────────────────────
  function getStoredTheme() {
    // 1. URL query parameter (highest priority for cross-page & file:// support)
    try {
      if (window.location && window.location.search) {
        var urlTheme = new URLSearchParams(window.location.search).get('theme');
        if (urlTheme === 'dark' || urlTheme === 'light') {
          return urlTheme;
        }
      }
    } catch (e) {}

    // 2. localStorage
    try {
      if (window.localStorage) {
        var local = localStorage.getItem(THEME_KEY);
        if (local === 'dark' || local === 'light') {
          return local;
        }
      }
    } catch (e) {}

    // 3. sessionStorage
    try {
      if (window.sessionStorage) {
        var session = sessionStorage.getItem(THEME_KEY);
        if (session === 'dark' || session === 'light') {
          return session;
        }
      }
    } catch (e) {}

    // 4. Cookies
    try {
      if (document.cookie) {
        var match = document.cookie.match(new RegExp('(?:^|;\\s*)' + THEME_KEY + '=([^;]+)'));
        if (match && (match[1] === 'dark' || match[1] === 'light')) {
          return match[1];
        }
      }
    } catch (e) {}

    // 5. System preference
    try {
      if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        return 'dark';
      }
    } catch (e) {}

    return 'light';
  }

  // ── Update Internal Navigation Links ────────────────────────
  // Ensures links carry the theme parameter across page navigation,
  // guaranteeing 100% reliability even under file:// protocol.
  function updateInternalLinks(theme) {
    if (!theme) return;
    try {
      var links = document.querySelectorAll('a[href]');
      links.forEach(function (a) {
        var href = a.getAttribute('href');
        if (!href) return;
        // Skip external, hash-only, or special links
        if (
          href.startsWith('http://') ||
          href.startsWith('https://') ||
          href.startsWith('#') ||
          href.startsWith('mailto:') ||
          href.startsWith('javascript:')
        ) {
          return;
        }

        var parts = href.split('#');
        var base = parts[0];
        var hash = parts[1] ? '#' + parts[1] : '';

        var qParts = base.split('?');
        var path = qParts[0];
        var search = qParts[1] ? new URLSearchParams(qParts[1]) : new URLSearchParams();
        search.set('theme', theme);

        a.setAttribute('href', path + '?' + search.toString() + hash);
      });
    } catch (e) {}
  }

  // ── Multi-tier Theme Saving ─────────────────────────────────
  function setStoredTheme(theme) {
    // 1. localStorage
    try {
      if (window.localStorage) {
        localStorage.setItem(THEME_KEY, theme);
      }
    } catch (e) {}

    // 2. sessionStorage
    try {
      if (window.sessionStorage) {
        sessionStorage.setItem(THEME_KEY, theme);
      }
    } catch (e) {}

    // 3. Cookie (valid for 1 year)
    try {
      document.cookie = THEME_KEY + '=' + theme + '; path=/; max-age=31536000; SameSite=Lax';
    } catch (e) {}

    // 4. Update internal links
    updateInternalLinks(theme);

    // 5. Update current URL parameter without reload
    try {
      if (window.history && window.history.replaceState && window.location) {
        var url = new URL(window.location.href);
        url.searchParams.set('theme', theme);
        window.history.replaceState(null, '', url.toString());
      }
    } catch (e) {}
  }

  // ── Apply Theme to DOM ──────────────────────────────────────
  function applyTheme(theme, save) {
    if (theme !== 'dark' && theme !== 'light') {
      theme = 'light';
    }

    document.documentElement.setAttribute('data-theme', theme);

    if (save) {
      setStoredTheme(theme);
    } else {
      updateInternalLinks(theme);
    }

    // Update all theme toggle buttons across the page
    var buttons = document.querySelectorAll('.theme-toggle');
    buttons.forEach(function (btn) {
      btn.textContent = theme === 'dark' ? '☀️' : '🌙';
      var label = theme === 'dark' ? 'Switch to light theme' : 'Switch to dark theme';
      btn.setAttribute('aria-label', label);
      btn.setAttribute('title', label);
    });
  }

  // Apply resolved theme immediately
  var initialTheme = getStoredTheme();
  applyTheme(initialTheme, false);

  // ── DOM Ready Initializations ───────────────────────────────
  document.addEventListener('DOMContentLoaded', function () {
    // Re-apply to update button icons and internal links after DOM parse
    var currentTheme = document.documentElement.getAttribute('data-theme') || getStoredTheme();
    applyTheme(currentTheme, false);

    // ── Theme Toggle Event Listeners ──────────────────────────
    var toggleButtons = document.querySelectorAll('.theme-toggle');
    toggleButtons.forEach(function (btn) {
      btn.addEventListener('click', function (e) {
        e.preventDefault();
        var active = document.documentElement.getAttribute('data-theme') || 'light';
        var next = active === 'dark' ? 'light' : 'dark';
        applyTheme(next, true);
      });
    });

    // ── System Theme Preference Change Listener ───────────────
    if (window.matchMedia) {
      try {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function (e) {
          // Only switch if user hasn't explicitly saved a choice in localStorage/sessionStorage
          var hasExplicitChoice = false;
          try {
            hasExplicitChoice = !!(localStorage.getItem(THEME_KEY) || sessionStorage.getItem(THEME_KEY));
          } catch (err) {}

          if (!hasExplicitChoice) {
            applyTheme(e.matches ? 'dark' : 'light', false);
          }
        });
      } catch (err) {}
    }

    // ── Mobile Navigation ──────────────────────────────────────
    var hamburger = document.querySelector('.nav__hamburger');
    var navLinks = document.querySelector('.nav__links');

    if (hamburger && navLinks) {
      hamburger.addEventListener('click', function () {
        navLinks.classList.toggle('nav__links--open');
        hamburger.classList.toggle('nav__hamburger--active');
      });

      // Close mobile nav when clicking any nav link
      navLinks.querySelectorAll('.nav__link').forEach(function (link) {
        link.addEventListener('click', function () {
          navLinks.classList.remove('nav__links--open');
          hamburger.classList.remove('nav__hamburger--active');
        });
      });
    }

    // ── Install Modal ──────────────────────────────────────────
    var modalOverlay = document.querySelector('.modal-overlay');

    if (modalOverlay) {
      // Open modal
      document.querySelectorAll('[data-modal="install"]').forEach(function (trigger) {
        trigger.addEventListener('click', function (e) {
          e.preventDefault();
          modalOverlay.classList.add('modal-overlay--active');
          document.body.style.overflow = 'hidden';
        });
      });

      // Close modal
      function closeModal() {
        modalOverlay.classList.remove('modal-overlay--active');
        document.body.style.overflow = '';
      }

      var redDot = modalOverlay.querySelector('.modal__dot--red');
      if (redDot) redDot.addEventListener('click', closeModal);

      modalOverlay.querySelectorAll('[data-close-modal]').forEach(function (el) {
        el.addEventListener('click', closeModal);
      });

      modalOverlay.addEventListener('click', function (e) {
        if (e.target === modalOverlay) closeModal();
      });

      // Close on Escape key
      document.addEventListener('keydown', function (e) {
        if (e.key === 'Escape' && modalOverlay.classList.contains('modal-overlay--active')) {
          closeModal();
        }
      });
    }

    // ── Copy to Clipboard ──────────────────────────────────────
    document.querySelectorAll('[data-copy]').forEach(function (btn) {
      btn.addEventListener('click', function () {
        var text = btn.getAttribute('data-copy');
        if (!text) {
          var codeBlock = btn.closest('.code-block');
          if (codeBlock) {
            text = codeBlock.querySelector('code')?.textContent || '';
          }
          var installCmd = btn.closest('.install-cmd');
          if (installCmd) {
            text = installCmd.querySelector('.install-cmd__text')?.textContent || '';
          }
        }

        if (text) {
          navigator.clipboard.writeText(text.trim()).then(function () {
            var original = btn.textContent;
            btn.textContent = 'Copied!';
            setTimeout(function () {
              btn.textContent = original;
            }, 1500);
          }).catch(function () {
            // Fallback for older browsers or restricted permissions
            var textArea = document.createElement('textarea');
            textArea.value = text.trim();
            document.body.appendChild(textArea);
            textArea.select();
            try {
              document.execCommand('copy');
              var orig = btn.textContent;
              btn.textContent = 'Copied!';
              setTimeout(function () {
                btn.textContent = orig;
              }, 1500);
            } catch (err) {}
            document.body.removeChild(textArea);
          });
        }
      });
    });

    // ── Active Nav Link ────────────────────────────────────────
    try {
      var path = window.location.pathname;
      document.querySelectorAll('.nav__link').forEach(function (link) {
        var href = link.getAttribute('href');
        if (!href) return;

        // Clean href of queries/hashes
        var cleanHref = href.split('?')[0].split('#')[0];
        var linkPath = cleanHref.replace(/\/index\.html$/, '/').replace(/\/$/, '') || '/';
        var currentPath = path.replace(/\/index\.html$/, '/').replace(/\/$/, '') || '/';

        if (currentPath === linkPath || (linkPath !== '/' && currentPath.endsWith(linkPath))) {
          link.classList.add('nav__link--active');
        }
      });
    } catch (e) {}
  });
})();
