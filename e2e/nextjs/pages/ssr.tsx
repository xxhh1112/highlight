import { useRouter } from 'next/router'

type Props = {
	date: string
	random: number
}
export default function SsrPage({ date, random }: Props) {
	const router = useRouter()
	const isError = router.asPath.includes('error')

	if (isError) {
		throw new Error('🎉 SSR Error: src/pages/ssr.tsx')
	}

	return (
		<div>
			<h1>SSR Lives</h1>
			<p>The random number is {random}</p>
			<p>The date is {date}</p>
		</div>
	)
}

export async function getStaticProps() {
	console.info('getStaticProps pages/ssr')

	return {
		props: {
			random: Math.random(),
			date: new Date().toISOString(),
		},
		revalidate: 10, // seconds
	}
}
