<script lang="ts">
	let { size = 'md' }: { size?: 'sm' | 'md' | 'lg' } = $props();
	
	const sizes = {
		sm: { logo: 24, text: 14 },
		md: { logo: 32, text: 18 },
		lg: { logo: 40, text: 22 }
	};
</script>

<div class="ekko-logo" style="--logo-size: {sizes[size].logo}px; --text-size: {sizes[size].text}px;">
	<div class="logo-circle">
		<svg viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
			<circle cx="16" cy="16" r="16" fill="url(#gradient)"/>
			<path d="M10 16C10 12.6863 12.6863 10 16 10C19.3137 10 22 12.6863 22 16" stroke="white" stroke-width="2.5" stroke-linecap="round"/>
			<path d="M13 16C13 14.3431 14.3431 13 16 13C17.6569 13 19 14.3431 19 16" stroke="white" stroke-width="2" stroke-linecap="round"/>
			<defs>
				<linearGradient id="gradient" x1="0" y1="0" x2="32" y2="32">
					<stop offset="0%" stop-color="#22c55e"/>
					<stop offset="100%" stop-color="#0ea5e9"/>
				</linearGradient>
			</defs>
		</svg>
	</div>
	<span class="logo-text">e<span class="k-flip">k</span>ko</span>
</div>

<style>
	.ekko-logo {
		display: flex;
		align-items: center;
		gap: 6px;
	}

	.logo-circle {
		width: var(--logo-size);
		height: var(--logo-size);
	}

	.logo-circle svg {
		width: 100%;
		height: 100%;
	}

	.logo-text {
		font-size: var(--text-size);
		font-weight: 700;
		color: var(--color-primary-dark);
		letter-spacing: -0.02em;
	}

	.k-flip {
		display: inline-block;
		transform: scaleX(-1);
	}
</style>
