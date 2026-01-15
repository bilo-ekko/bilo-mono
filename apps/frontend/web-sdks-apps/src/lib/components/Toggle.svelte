<script lang="ts">
	let { checked = $bindable(false), onchange }: { checked: boolean; onchange?: (value: boolean) => void } = $props();
	
	function handleClick() {
		checked = !checked;
		onchange?.(checked);
	}
</script>

<button
	class="toggle"
	class:active={checked}
	onclick={handleClick}
	role="switch"
	aria-checked={checked}
>
	<span class="toggle-thumb"></span>
</button>

<style>
	.toggle {
		position: relative;
		width: 52px;
		height: 28px;
		background: var(--color-toggle-bg);
		border-radius: 14px;
		transition: background-color 0.2s ease;
		flex-shrink: 0;
	}

	.toggle.active {
		background: var(--color-toggle-active);
	}

	.toggle-thumb {
		position: absolute;
		top: 2px;
		left: 2px;
		width: 24px;
		height: 24px;
		background: white;
		border-radius: 50%;
		transition: transform 0.2s ease;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
	}

	.toggle.active .toggle-thumb {
		transform: translateX(24px);
	}
</style>
