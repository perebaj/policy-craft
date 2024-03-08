import { PlusCircleIcon } from '@heroicons/react/20/solid'
import { useContext } from 'react'
import {
  BaseEdge,
  EdgeLabelRenderer,
  EdgeProps,
  getSmoothStepPath,
} from 'reactflow'

import { DrawerName, editor } from './Editor'

export function AddNodeEdge({
  id,
  sourceX,
  sourceY,
  targetX,
  targetY,
  sourcePosition,
  targetPosition,
  markerEnd,
  label,
  source,
  target,
}: EdgeProps) {
  const { showDrawer } = useContext(editor)
  const [edgePath, centerX, centerY] = getSmoothStepPath({
    sourceX,
    sourceY,
    sourcePosition,
    targetX,
    targetY,
    targetPosition,
  })

  const onEdgeClick = (
    event: Record<'stopPropagation', () => void>,
    id: string,
  ) => {
    event.stopPropagation()
    showDrawer(DrawerName.newNode, {
      id,
      sourceNodeId: source,
      targetNodeId: target,
      sourceEdgeLabel: label,
    })
  }

  return (
    <>
      <BaseEdge
        path={edgePath}
        markerEnd={markerEnd}
        style={{ pointerEvents: 'none', strokeWidth: 2 }}
      />
      <EdgeLabelRenderer>
        <div
          style={{
            transform: `translate(-50%, -50%) translate(${centerX}px,${centerY}px)`,
            pointerEvents: 'all',
          }}
          className="absolute h-4 w-4 rounded-full text-black"
        >
          <button
            onClick={(event) => onEdgeClick(event, id)}
            className={`shadow-sm-light group absolute flex h-full w-full items-center justify-center overflow-hidden rounded-full border border-teal-400 shadow-gray-300 transition-colors duration-300 hover-focus:border-slate-900 hover-focus:bg-black`}
          >
            <PlusCircleIcon className="absolute w-[130%] bg-white text-teal-500 transition-colors duration-300 group-hover-focus:bg-slate-900 group-hover-focus:text-teal-300" />
          </button>
        </div>
      </EdgeLabelRenderer>
    </>
  )
}
