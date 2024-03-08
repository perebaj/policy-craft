import { ReactNode } from 'react'

type ChooseNodeButtonProps = {
  label: string
  preview: ReactNode
  onClick: () => void
}

export const ChooseNodeButton = ({
  label,
  preview,
  onClick,
}: ChooseNodeButtonProps) => (
  <button
    className="group flex cursor-pointer flex-col items-center justify-center rounded-md border border-N-500 bg-white p-2 shadow-small shadow-neutral-300 hover-focus:bg-N-300"
    onClick={onClick}
  >
    <div className="flex h-20 w-20 items-center">{preview}</div>
    <div className="font-lato mt-3 flex text-center font-medium group-hover-focus:font-medium">
      <p>{label}</p>
    </div>
  </button>
)
