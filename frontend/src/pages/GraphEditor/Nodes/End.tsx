import { Handle, NodeProps, Position } from 'reactflow'

import { NodeWrapper } from './NodeWrapper'

type EndNodeData = {
  width: number
  height: number
}

export function EndNode({ data }: NodeProps<EndNodeData>) {
  return (
    <NodeWrapper>
      <div
        className={`flex aspect-square h-full cursor-pointer items-center justify-center rounded-full border-4 border-N-400 bg-white text-[12px]`}
        style={{
          width: data.width,
          height: data.height,
        }}
      >
        <Handle
          type="target"
          id="target"
          className="invisible"
          position={Position.Top}
          isConnectable={false}
        />
        <p className="font-medium">{'End'}</p>
      </div>
    </NodeWrapper>
  )
}
