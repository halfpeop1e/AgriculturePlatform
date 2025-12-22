import type { Expert, ExpertDetail } from './expert'

const baseIntro =
	'Seasoned agricultural consultant with a focus on precision farming and sustainable practices.'

const makeExpert = (overrides: Partial<Expert>): Expert => ({
	id: `expert_${Math.random().toString(36).slice(2, 8)}`,
	name: 'Demo Specialist',
	avatar: 'https://picsum.photos/seed/expert/120',
	field: 'Smart Agriculture',
	introduction: baseIntro,
	skills: ['Crop Rotation', 'IoT Monitoring', 'Soil Analysis'],
	consultCount: 128,
	rating: 4.8,
	responseTime: 'Within 24h',
	...overrides
})

export const mockExperts: Expert[] = [
	makeExpert({
		name: 'Alice Chen',
		field: 'Soil Science',
		skills: ['Soil Testing', 'Nutrient Planning', 'Microbiome Tuning'],
		avatar: 'https://picsum.photos/seed/expert-1/120',
		consultCount: 205,
		rating: 4.9
	}),
	makeExpert({
		name: 'Brian Lee',
		field: 'Horticulture',
		skills: ['Greenhouse Control', 'Hydroponics', 'Pest Management'],
		avatar: 'https://picsum.photos/seed/expert-2/120',
		consultCount: 162,
		rating: 4.7
	}),
	makeExpert({
		name: 'Chiara Wu',
		field: 'Livestock Management',
		skills: ['Feed Optimization', 'Health Monitoring', 'Breeding Strategy'],
		avatar: 'https://picsum.photos/seed/expert-3/120',
		consultCount: 94,
		rating: 4.6
	}),
	makeExpert({
		name: 'Daniel Kim',
		field: 'Agri-Tech',
		skills: ['Drone Survey', 'Yield Forecast', 'Automation'],
		avatar: 'https://picsum.photos/seed/expert-4/120',
		consultCount: 140,
		rating: 4.8
	})
]

export const mockExpertDetails: ExpertDetail[] = mockExperts.map((expert, index) => ({
	...expert,
	education: index % 2 === 0 ? 'PhD in Agronomy, Midwest University' : 'MSc in Precision Farming, Tech College',
	experience: `${8 + index} years in field advisory roles focusing on data-driven agriculture.`,
	certification: ['ISO 14001 Auditor', 'Certified Crop Adviser'],
	availableTime: ['Mon-Fri 09:00-18:00', 'Sat 10:00-15:00'],
	serviceScope: ['Remote Consulting', 'On-site Assessment', 'Training Sessions'],
	price: 299 + index * 20,
	recentCases: [
		{
			date: '2024-08-12',
			question: 'How to improve soil moisture retention in sandy fields?',
			answer: 'Recommended biochar application combined with cover crops and drip irrigation.'
		},
		{
			date: '2024-05-03',
			question: 'Livestock feed plan for high-yield dairy herd.',
			answer: 'Designed a phased nutrition plan with digestibility monitoring.'
		}
	]
}))

export const getMockExpertById = (id: string): ExpertDetail | undefined =>
	mockExpertDetails.find((expert) => expert.id === id)
