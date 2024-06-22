import axiosInstance from './axios/instance'
import ApiGenerator from '@/fetch/apiGenerator/index'

const swaggerContext = import.meta.glob('../../.swagger/*.swagger.json')

const apis = await ApiGenerator.createAllAPI(swaggerContext)
const fetch = ApiGenerator.createFetch(apis, axiosInstance)

if (process.env.NODE_ENV === 'development') {
  console.table(ApiGenerator.collectSwaggerInfo(swaggerContext))
}

/**
 * [Swagger Parameter Types]:
 * - Path: path parameters, such as /users/{id}
 * - Body: request body, such as POST/PUT data
 * - Query: query parameters, such as /users?role=admin
 * - Header: header parameters, such as X-MyHeader: Value
 */
export default fetch as (
  operationId: string,
  data?: Partial<{
    path: object
    body: any
    query: object
    header: object
  }>
) => Promise<any>
