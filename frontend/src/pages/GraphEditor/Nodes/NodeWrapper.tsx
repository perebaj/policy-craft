import { memo, PropsWithChildren } from 'react'

const BaseWrapper = ({ children }: PropsWithChildren) => {
  return (
    <div
      className={`group/container font-lato flex flex-col items-center justify-center`}
    >
      {children}
    </div>
  )
}

export const NodeWrapper = memo(BaseWrapper)
