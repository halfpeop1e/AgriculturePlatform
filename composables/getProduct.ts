/*
  getProduct.ts
  说明：获取产品列表的网络请求封装
    参数：options (可选) — { page?: number; pageSize?: number; salerId?: string; onlyMine?: boolean }
  返回：成功时返回标准化后的 ProductListResult，包含 list、total、page、pageSize、hasMore
  依赖：内部使用 useAxios() 创建的 Axios 实例（会自动注入 token 等拦截逻辑）
  注意：后端返回的数据包含分页信息与原始字段，需要在此处进行格式转换。
*/
import { useAxios } from "./useAxios"
import { useUserStore } from "~/utils/userStore"
import type { ProductListPayload, ProductListResult } from "~/types/product";

const useAxiosInstance = useAxios()

export interface GetProductListOptions {
    page?: number
    pageSize?: number
    salerId?: string
    onlyMine?: boolean
}

interface ProductListApiResponse {
    code: number
    message: string
    data?: ProductListPayload
}

export async function getProductList(options: GetProductListOptions = {}): Promise<ProductListResult> {
    const { page = 1, pageSize = 10, salerId, onlyMine = false } = options
    const userStore = useUserStore()
    const currentUserId = userStore.userId
    const effectiveSalerId = onlyMine ? currentUserId : salerId
    console.log(options)

    if (onlyMine && !currentUserId) {
        console.warn('请求个人商品列表时未找到用户 ID')
        return {
            list: [],
            total: 0,
            page,
            pageSize,
            hasMore: false,
        }
    }
    try {
        const params: Record<string, unknown> = {
            page,
            pageSize,
        }

        if (effectiveSalerId && effectiveSalerId !== 'all') {
            params.salerId = effectiveSalerId
        }

        const response = await useAxiosInstance.get<ProductListApiResponse>("/product/list", {
            params
        })

        if (response.status === 200 && response.data?.code === 200 && response.data.data) {
            const payload = response.data.data
            const normalizedList = (payload.productList || []).map((item) => ({
                id: String(item.id),
                name: item.name,
                image: Array.isArray(item.image) ? item.image.map((url) => ({ url })) : [],
                description: item.description,
                price: item.price,
                stock: Number(item.stock),
                saler: item.saler,
                salerId: item.salerId,
                category: item.category,
            }))

            const filteredList = (effectiveSalerId && effectiveSalerId !== 'all')
                ? normalizedList.filter((item) => item.salerId === effectiveSalerId)
                : normalizedList

            const totalFromServer = typeof payload.total === 'number' ? payload.total : normalizedList.length
            const total = (effectiveSalerId && effectiveSalerId !== 'all')
                ? filteredList.length
                : totalFromServer

            const currentPage = typeof payload.page === 'number' ? payload.page : page
            const currentPageSize = typeof payload.pageSize === 'number' ? payload.pageSize : pageSize
            const computedHasMore = (effectiveSalerId && effectiveSalerId !== 'all')
                ? false
                : typeof payload.hasMore === 'boolean'
                    ? payload.hasMore
                    : currentPage * currentPageSize < total

            const result: ProductListResult = {
                list: filteredList,
                total,
                page: currentPage,
                pageSize: currentPageSize,
                hasMore: computedHasMore,
            }

            console.log('获取产品列表成功', result)

            return result
        }

        throw new Error(response.data?.message || '获取产品列表失败')
    }
    catch(err){
        console.error('获取产品列表失败',err)
        // 将错误向上抛出，调用方负责重试/处理逻辑
        throw err
    }
}
