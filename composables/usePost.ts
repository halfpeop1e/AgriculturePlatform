import { ref } from 'vue'
import type { Post, Reply, Author } from '~/types/post'
import { useAxios } from './useAxios'
import { mockPosts } from '~/types/posttest'

export interface CreatePostInput {
	title: string
	content: string
	author: Author
}

export interface CreateReplyInput {
	content: string
	author: Author
}

const axiosInstance = useAxios()

const postsState = ref<Post[]>([])
const loadingState = ref(false)
const errorState = ref<string | null>(null)

const upsertPost = (post: Post) => {
	const idx = postsState.value.findIndex((item) => item.id === post.id)
	if (idx >= 0) {
		postsState.value[idx] = { ...postsState.value[idx], ...post }
	} else {
		postsState.value.unshift(post)
	}
}

const normalizePost = (post: Post): Post => ({
	...post,
	likes: post.likes ?? 0,
	liked: post.liked ?? false,
	replies: post.replies?.map((reply) => normalizeReply(reply)) ?? []
})

const normalizeReply = (reply: Reply): Reply => ({
	...reply,
	author: reply.author
		? {
			id: reply.author.id,
			nickname: reply.author.nickname || '匿名用户',
			avatar: reply.author.avatar
		}
		: {
			id: 'anonymous',
			nickname: '匿名用户',
			avatar: ''
		},
	likes: reply.likes ?? 0,
	liked: reply.liked ?? false
})

export const usePost = () => {
	const setError = (message: string | null) => {
		errorState.value = message
	}

	const fetchPosts = async () => {
		loadingState.value = true
		setError(null)
		try {
			const response = await axiosInstance.get('/posts')
			const data = response.data.data || []
			postsState.value = data.map(normalizePost)
			return postsState.value
		} catch (err: any) {
			console.error('获取帖子列表失败，回退到 mock 数据', err)
			postsState.value = mockPosts.map(normalizePost)
			setError(err?.message || '获取帖子列表失败，已显示测试数据')
			return postsState.value
		} finally {
			loadingState.value = false
		}
	}

	const fetchPostById = async (id: string) => {
		const existing = postsState.value.find((post) => post.id === id)
		if (existing) {
			console.log('使用已存在帖子数据', existing)
			return existing
		}

		loadingState.value = true
		setError(null)
		try {
			const response = await axiosInstance.get(`/posts/${id}`)
			if (response.data) {
				const normalized = normalizePost(response.data.data)
				upsertPost(normalized)
				return normalized
			}
			return null
		} catch (err: any) {
			console.error('获取帖子详情失败，尝试使用 mock 数据', err)
			const mockPost = mockPosts.find((item) => item.id === id)
			if (mockPost) {
				const normalized = normalizePost(mockPost)
				upsertPost(normalized)
				setError('已加载测试帖子数据')
				return normalized
			}
			setError(err?.message || '获取帖子详情失败')
			throw err
		} finally {
			loadingState.value = false
		}
	}

	const createPost = async (input: CreatePostInput) => {
		const tempPost: Post = normalizePost({
			id: `post_${Date.now()}_${Math.random().toString(36).slice(2, 8)}`,
			title: input.title,
			content: input.content,
			author: input.author,
			likes: 0,
			liked: false,
			replies: [],
			createdAt: new Date().toISOString()
		})

		try {
			const response = await axiosInstance.post<Post>('/posts', tempPost)
			const created = response.data ? normalizePost(response.data) : tempPost
			upsertPost(created)
      ElMessage.success('帖子创建成功')
			return created
		} catch (err: any) {
			console.error('创建帖子失败，使用本地结果', err)
			upsertPost(tempPost)
			setError(err?.message || '创建帖子失败')
			return tempPost
		}
	}

	const addReply = async (postId: string, input: CreateReplyInput) => {
		const post = postsState.value.find((item) => item.id === postId)
		if (!post) {
			throw new Error('找不到对应帖子')
		}

		const tempReply: Reply = normalizeReply({
			id: `reply_${Date.now()}_${Math.random().toString(36).slice(2, 8)}`,
			content: input.content,
			author: input.author,
			likes: 0,
			liked: false,
			createdAt: new Date().toISOString()
		})

		try {
			const response = await axiosInstance.post(`/posts/${postId}/replies`, tempReply)
	      console.log('回复帖子成功', response.data)
			const createdPayload = response.data?.data ?? response.data
			const created = createdPayload ? normalizeReply(createdPayload) : tempReply
			post.replies = [created, ...(post.replies || [])]
	      ElMessage.success('回复帖子成功')
	      console.log(created, post.replies)
			return created
		} catch (err: any) {
			console.error('回复帖子失败，使用本地结果', err)
			post.replies = [tempReply, ...(post.replies || [])]
			setError(err?.message || '回复帖子失败')
			return tempReply
		}
	}

	const togglePostLike = async (postId: string, liked?: boolean) => {
		const post = postsState.value.find((item) => item.id === postId)
		if (!post) return

		const nextLiked = liked ?? !post.liked
		const delta = nextLiked ? 1 : -1
		const original = { liked: post.liked, likes: post.likes ?? 0 }

		post.liked = nextLiked
		post.likes = Math.max(0, (post.likes ?? 0) + delta)

		try {
			await axiosInstance.post(`/posts/${postId}/like`, { liked: nextLiked })
		} catch (err: any) {
			console.error('更新帖子点赞状态失败，回滚', err)
      ElMessage.error('点赞失败')
			post.liked = original.liked
			post.likes = original.likes
			setError(err?.message || '更新点赞状态失败')
		}
	}

	const toggleReplyLike = async (postId: string, replyId: string, liked?: boolean) => {
		const post = postsState.value.find((item) => item.id === postId)
		if (!post) return
		const reply = post.replies?.find((item) => item.id === replyId)
		if (!reply) return

		const nextLiked = liked ?? !reply.liked
		const delta = nextLiked ? 1 : -1
		const original = { liked: reply.liked, likes: reply.likes ?? 0 }

		reply.liked = nextLiked
		reply.likes = Math.max(0, (reply.likes ?? 0) + delta)

		try {
			await axiosInstance.post(`/posts/${postId}/replies/${replyId}/like`, { liked: nextLiked })
		} catch (err: any) {
			console.error('更新回复点赞状态失败，回滚', err)
			reply.liked = original.liked
			reply.likes = original.likes
			setError(err?.message || '更新回复点赞状态失败')
		}
	}

	const getPostById = (id: string) => postsState.value.find((post) => post.id === id) || null

	return {
		posts: postsState,
		loading: loadingState,
		error: errorState,
		fetchPosts,
		fetchPostById,
		createPost,
		addReply,
		togglePostLike,
		toggleReplyLike,
		getPostById
	}
}

export default usePost
