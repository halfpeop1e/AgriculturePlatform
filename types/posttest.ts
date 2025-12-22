import type { Author, Post, Reply } from './post'

const baseAuthor: Author = {
	id: 'author_1',
	nickname: 'Demo User',
	avatar: 'https://picsum.photos/seed/post-author/80'
}

const makeReply = (overrides: Partial<Reply>): Reply => ({
	id: `reply_${Math.random().toString(36).slice(2, 8)}`,
	content: 'Sample reply copied to verify layout and interaction of the answer card.',
	author: {
		id: 'reply_author_1',
		nickname: 'Responder',
		avatar: 'https://picsum.photos/seed/post-reply/64'
	},
	likes: 2,
	liked: false,
	createdAt: new Date().toISOString(),
	...overrides
})

const makePost = (overrides: Partial<Post>): Post => ({
	id: `post_${Math.random().toString(36).slice(2, 8)}`,
	title: '示例帖子标题',
	content:
		'Example body text for the post card. Validate spacing, line breaks, and long paragraph rendering.\nLine breaks should be preserved.',
	author: baseAuthor,
	likes: 12,
	liked: false,
	replies: [
		makeReply({
			author: {
				id: 'reply_author_2',
				nickname: 'Alice',
				avatar: 'https://picsum.photos/seed/reply-1/64'
			},
			likes: 5
		}),
		makeReply({
			author: {
				id: 'reply_author_3',
				nickname: 'Bob',
				avatar: 'https://picsum.photos/seed/reply-2/64'
			},
			content: 'Another reply that helps validate multiple entries and like-state toggling.'
		})
	],
	createdAt: new Date().toISOString(),
	...overrides
})

export const mockPosts: Post[] = [
	makePost({
		title: 'Wheat Management Checklist',
		content:
			'A concise summary of sowing, fertilising, and irrigation checkpoints to test PostCard typography.',
		likes: 18
	}),
	makePost({
		title: 'Smart Irrigation Takeaways',
		liked: true,
		likes: 47,
		replies: [
			makeReply({
				content: 'Looking forward to a hardware breakdown and more configuration tips!',
				likes: 8,
				liked: true
			})
		]
	}),
	makePost({
		title: 'Produce Livestream Notes',
		content: 'Collected observations on running livestream sessions and solving common issues.',
		replies: []
	})
]

export const getMockPostById = (id: string): Post | undefined =>
	mockPosts.find((post) => post.id === id)

export const toggleMockLike = (post: Post): Post => ({
	...post,
	liked: !post.liked,
	likes: Math.max(0, (post.likes ?? 0) + (post.liked ? -1 : 1))
})

export const addMockReply = (post: Post, replyOverrides: Partial<Reply>) => {
	const reply = makeReply(replyOverrides)
	post.replies = [reply, ...(post.replies || [])]
	return reply
}
