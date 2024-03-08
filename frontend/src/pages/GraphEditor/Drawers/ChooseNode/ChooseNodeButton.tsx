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
    className="border-N-500 shadow-small hover-focus:bg-N-300 group flex cursor-pointer flex-col items-center justify-center rounded-md border bg-white p-2 shadow-neutral-300"
    onClick={onClick}
  >
    <div className="flex h-20 w-20 items-center">{preview}</div>
    <div className="font-lato group-hover-focus:font-medium mt-3 flex text-center font-medium">
      <p>{label}</p>
    </div>
  </button>
)
