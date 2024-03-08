import { ChevronDoubleRightIcon } from '@heroicons/react/20/solid'
import { motion } from 'framer-motion'
import { memo, ReactNode } from 'react'

type DrawerProps = {
  title: string
  content: ReactNode
  visible: boolean
  onClose?: () => void
}
// Arbitrary value.

export const DRAWER_ANIMATION_IN_SECONDS = 0.3
export const DRAWER_ANIMATION_IN_MILLISECONDS =
  DRAWER_ANIMATION_IN_SECONDS * 1000

/**
 * This is a base and generic drawer component that is customizable based on props passed
 * @param props Drawer properties: title, content, visible...
 * @returns A customized Drawer component
 */
const CustomDrawer = ({ title, content, onClose, visible }: DrawerProps) => {
  if (!visible) {
    return null
  }

  return (
    <motion.div
      className={`border-border-normal fixed right-0 z-40 h-full w-2/5 overflow-y-auto border-l-2 bg-N-100 p-4`}
      animate={{
        x: visible ? '0' : '100%',
      }}
      initial={{
        x: '100%',
      }}
      transition={{
        bounce: 0,
        duration: DRAWER_ANIMATION_IN_SECONDS,
      }}
      aria-labelledby="drawer-title"
      data-testid="drawer"
    >
      <div className="mb-4 flex items-center">
        <h3
          id="drawer-title"
          className="text-text-headline font-lato  inline-flex w-full items-center text-lg font-semibold"
        >
          {title}
        </h3>
        <button
          type="button"
          data-drawer-hide="drawer-close-button"
          aria-controls="drawer-close-button"
          onClick={onClose}
          className="absolute right-2.5 flex h-8 w-8 items-center justify-center rounded-full text-gray-500 hover-focus:bg-gray-200 hover-focus:text-black"
          tabIndex={visible ? 0 : -1}
        >
          <ChevronDoubleRightIcon className="w-6" />
          <span className="sr-only">Close menu</span>
        </button>
      </div>
      {visible && content}
    </motion.div>
  )
}

export const Drawer = memo(CustomDrawer)
