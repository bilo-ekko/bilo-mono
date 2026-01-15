<script lang="ts">
	const partners = [
		{ name: 'Conservation International', abbr: 'CI', color: '#1a5632' },
		{ name: 'RSPB', abbr: 'RSPB', color: '#2d5a27' },
		{ name: 'Tusk', abbr: 'Tusk', color: '#d4432c' },
		{ name: '1% for the Planet', abbr: '1%', color: '#006eb6' },
		{ name: 'Conservation Collective', abbr: 'CC', color: '#2d7a3f' },
		{ name: 'Justdiggit', abbr: 'JD', color: '#7cb342' }
	];
</script>

<div class="partners">
	{#each partners as partner}
		<div class="partner-logo" style="--partner-color: {partner.color}" title={partner.name}>
			<span class="partner-abbr">{partner.abbr}</span>
		</div>
	{/each}
</div>

<style>
	.partners {
		display: flex;
		align-items: center;
		gap: 6px;
		flex-wrap: wrap;
	}

	.partner-logo {
		width: 28px;
		height: 28px;
		border-radius: 50%;
		background: var(--partner-color);
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 8px;
		font-weight: 700;
		color: white;
		text-transform: uppercase;
		letter-spacing: -0.02em;
	}

	.partner-abbr {
		line-height: 1;
	}
</style>
