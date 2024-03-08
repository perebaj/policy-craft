import { DiamondSvg } from 'assets/Diamond'
import { Handle, NodeProps, Position } from 'reactflow'

import { NodeWrapper } from './NodeWrapper'

export type ConditionalNodeData = {
  label: string
  elseLabel?: string
  width: number
  height: number
}

export function ConditionalNode({ data }: NodeProps<ConditionalNodeData>) {
  return (
    <NodeWrapper>
      <div
        className="group h-full w-full"
        style={{
          width: data.width,
          height: data.height,
        }}
      >
        <div className="relative flex h-full w-full  items-center justify-center p-9 text-center text-[12px]">
          <div
            className={`absolute left-0 top-0 z-0 h-full w-full text-Y-300 group-hover:text-Y-350 group-hover-focus:cursor-pointer [&>svg]:stroke-Y-600`}
          >
            <DiamondSvg strokeWidth={4} />
          </div>
          <Handle
            type="target"
            id="target"
            className="invisible"
            position={Position.Top}
            isConnectable={false}
          />
          <p className={`z-10 line-clamp-3 cursor-pointer`}>{data.label}</p>
          <Handle
            type="source"
            id="source"
            className="invisible"
            position={Position.Bottom}
            isConnectable={false}
          />
        </div>
      </div>
    </NodeWrapper>
  )
}
