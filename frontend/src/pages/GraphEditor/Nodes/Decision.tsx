import { Handle, NodeProps, Position } from 'reactflow'

import { NodeWrapper } from './NodeWrapper'

type DecisionNodeData = {
  width: number
  height: number
}

export function DecisionNode({ data }: NodeProps<DecisionNodeData>) {
  return (
    <NodeWrapper>
      <div
        className={`flex aspect-square h-full cursor-pointer items-center justify-center rounded-full border-4 border-red-300 bg-white text-[12px]`}
        style={{
          width: data.width,
          height: data.height,
        }}
      >
        <Handle
          type="target"
          id="decision"
          className="invisible"
          position={Position.Top}
          isConnectable={false}
        />
        <p className="font-medium">{'False'}</p>
      </div>
    </NodeWrapper>
  )
}
