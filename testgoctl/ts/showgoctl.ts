import webapi from "./gocliRequest"
import * as components from "./showgoctlComponents"
export * from "./showgoctlComponents"

/**
 * @description ping server
 * @param req
 */
export function pinghandler(req: components.PRequist) {
	return webapi.post<components.PResponse>(`/ping`, req)
}
