import { setupWorker } from 'msw/browser'
import { handlers } from './handlers'

export const enableMocking = async () => {
  if (process.env.NODE_ENV === 'development') {
    const worker = setupWorker(...handlers)
    await worker.start()
  }
}
