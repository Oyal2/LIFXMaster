import type { NavigateFunction } from 'react-router-dom'

export const backClick = (navigate: NavigateFunction, link: string): void => {
  navigate(link)
}
