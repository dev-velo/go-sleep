<script>
  import { onMount } from 'svelte';
  import * as SleepService from '../bindings/changeme/sleepservice.js';

  /**
   * @typedef {import('../bindings/changeme/models.js').AppInfo} AppInfo
   * @typedef {import('../bindings/changeme/models.js').DashboardSummary} DashboardSummary
   */

  let appInfo = $state(/** @type {AppInfo} */ ({
    name: 'Go Sleep',
    version: '',
    description: 'A local-first CPAP and sleep data dashboard.'
  }));

  let summary = $state(/** @type {DashboardSummary | null} */ (null));
  let errorMessage = $state('');
  let isLoading = $state(false);

  const cards = $derived(summary ? [
    { label: 'Last night usage', value: summary.lastNightUsage },
    { label: 'AHI', value: summary.ahi.toFixed(1) },
    { label: 'Leak 95%', value: `${summary.leak95.toFixed(1)} L/min` },
    { label: 'Pressure 95%', value: `${summary.pressure95.toFixed(1)} cmH2O` },
    { label: 'Imported nights', value: summary.importedNightCount }
  ] : []);

  async function loadAppInfo() {
    try {
      appInfo = await SleepService.AppInfo();
    } catch (error) {
      console.error('Failed to load app info:', error);
      errorMessage = 'Could not load app information from the backend.';
    }
  }

  async function refreshSummary() {
    isLoading = true;
    errorMessage = '';

    try {
      summary = await SleepService.GetDashboardSummary();
    } catch (error) {
      console.error('Failed to load dashboard summary:', error);
      errorMessage = 'Could not load the dashboard summary from the backend.';
    } finally {
      isLoading = false;
    }
  }

  onMount(() => {
    loadAppInfo();
    refreshSummary();
  });
</script>

<main class="app-shell">
  <header class="header">
    <div>
      <p class="eyebrow">Version {appInfo.version || '0.1.0'}</p>
      <h1>{appInfo.name}</h1>
      <p class="subtitle">{appInfo.description}</p>
    </div>

    <button class="refresh-button" onclick={refreshSummary} disabled={isLoading}>
      {isLoading ? 'Refreshing...' : 'Refresh Summary'}
    </button>
  </header>

  {#if errorMessage}
    <section class="error" role="alert">
      {errorMessage}
    </section>
  {/if}

  <section class="dashboard" aria-label="Dashboard summary">
    {#if summary}
      {#each cards as card}
        <article class="summary-card">
          <p>{card.label}</p>
          <strong>{card.value}</strong>
        </article>
      {/each}
    {:else}
      <article class="summary-card placeholder">
        <p>Loading dashboard summary...</p>
      </article>
    {/if}
  </section>

  {#if summary?.importedNightCount === 0}
    <section class="empty-state">
      <h2>No CPAP data imported yet.</h2>
      <p>Next step: add an SD card importer.</p>
    </section>
  {/if}
</main>

<style>
  .app-shell {
    min-height: 100vh;
    padding: 48px;
    color: #172033;
    background: #f7f8fb;
  }

  .header {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 24px;
    max-width: 1040px;
    margin: 0 auto 32px;
  }

  .eyebrow {
    margin: 0 0 8px;
    color: #5f6d7e;
    font-size: 0.85rem;
    font-weight: 700;
    letter-spacing: 0;
    text-transform: uppercase;
  }

  h1 {
    margin: 0;
    color: #101828;
    font-size: 2.75rem;
    line-height: 1.1;
  }

  .subtitle {
    max-width: 560px;
    margin: 12px 0 0;
    color: #475467;
    font-size: 1.05rem;
  }

  .refresh-button {
    min-width: 150px;
    border: 0;
    border-radius: 6px;
    padding: 12px 16px;
    color: white;
    background: #1f6feb;
    font: inherit;
    font-weight: 700;
    cursor: pointer;
  }

  .refresh-button:hover:not(:disabled) {
    background: #185abc;
  }

  .refresh-button:disabled {
    cursor: wait;
    opacity: 0.7;
  }

  .error {
    max-width: 1040px;
    margin: 0 auto 20px;
    border: 1px solid #f0b8b8;
    border-radius: 8px;
    padding: 14px 16px;
    color: #8a1f1f;
    background: #fff1f1;
  }

  .dashboard {
    display: grid;
    grid-template-columns: repeat(5, minmax(0, 1fr));
    gap: 16px;
    max-width: 1040px;
    margin: 0 auto;
  }

  .summary-card {
    min-height: 116px;
    border: 1px solid #d9e0ea;
    border-radius: 8px;
    padding: 18px;
    background: white;
    box-shadow: 0 1px 2px rgba(16, 24, 40, 0.06);
  }

  .summary-card p {
    margin: 0 0 14px;
    color: #667085;
    font-size: 0.9rem;
    font-weight: 700;
  }

  .summary-card strong {
    color: #101828;
    font-size: 1.75rem;
    line-height: 1.1;
  }

  .placeholder {
    grid-column: 1 / -1;
  }

  .empty-state {
    max-width: 1040px;
    margin: 24px auto 0;
    border: 1px dashed #b9c4d2;
    border-radius: 8px;
    padding: 20px;
    background: #ffffff;
  }

  .empty-state h2 {
    margin: 0 0 6px;
    color: #101828;
    font-size: 1.1rem;
  }

  .empty-state p {
    margin: 0;
    color: #475467;
  }

  @media (max-width: 860px) {
    .app-shell {
      padding: 28px 20px;
    }

    .header {
      flex-direction: column;
    }

    .refresh-button {
      width: 100%;
    }

    .dashboard {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  @media (max-width: 520px) {
    h1 {
      font-size: 2.2rem;
    }

    .dashboard {
      grid-template-columns: 1fr;
    }
  }
</style>
